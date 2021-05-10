package server

import (
	context "context"
	"event-service/dto"
	"event-service/filters"
	repo2 "event-service/interfaces/repository"
	service2 "event-service/interfaces/services"
	"event-service/pb"
	"event-service/repository"
	"event-service/services"
	"event-service/utils"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
)

type eventServer struct {
	repo repo2.EventRepository
	service service2.EventService
}

func (e *eventServer) CreateEvent(ctx context.Context, params *pb.EventInputParams) (*pb.EventResponse, error) {
	log.Println(fmt.Sprintf("getting creation request from client %v", params))
	err := utils.ValidateEventInput(params)
	if err != nil {
		return nil, err
	}
	// json parse
	bytes, err := params.Data.MarshalJSON()
	if err != nil {
		return nil, err
	}
	jsonData := datatypes.JSON(bytes)
	input := &dto.EventCreateParams{
		Email:       params.Email,
		Environment: params.Environment,
		Component:   params.Component,
		Message:     params.MessageString,
		Data:        jsonData,
	}
	event, err := e.service.CreateEvent(*input)
	if err != nil {
		return nil, err
	}
	return event.ConvertToMessage(), nil
}

func (e *eventServer) FilterEvents(ctx context.Context, input *pb.EventFilterInput) (*pb.EventFilterResponse, error) {
	eventFilter := filters.GetEventFilter(e.repo)
	log.Println(fmt.Sprintf("receiving request with params %v", input))
	if val, ok := input.Params["component"]; ok {
		eventFilter = eventFilter.OfComponent(&val)
	}
	if val, ok := input.Params["email"]; ok {
		eventFilter = eventFilter.Author(&val)
	}
	if val, ok := input.Params["environment"]; ok {
		eventFilter = eventFilter.InEnvironment(&val)
	}
	if val, ok := input.Params["message"]; ok {
		eventFilter = eventFilter.WithMessage(&val)
	}
	if val, ok := input.Params["from"]; ok {
		eventFilter = eventFilter.Since(&val)
	}



	events := eventFilter.Get()
	eventsMap := make([]*pb.EventResponse, len(events))

	for i, event := range events {
		eventsMap[i] = event.ConvertToMessage()
	}
	return &pb.EventFilterResponse{Events: eventsMap}, nil
}

func GetNewEventServer(db *gorm.DB) pb.EventServer {
	repo := repository.GetEventRepository(db)
	eventService := services.GetEventService(repo)
	return &eventServer{
		repo:    repo,
		service: eventService,
	}
}