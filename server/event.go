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

// physical server initialization
type eventServer struct {
	// repo interface
	repo repo2.EventRepository
	// service interface
	service service2.EventService
}

func (e *eventServer) CreateEvent(ctx context.Context, params *pb.EventInputParams) (*pb.EventResponse, error) {
	log.Println(fmt.Sprintf("getting creation request from client %v", params))
	// validation of input params
	err := utils.ValidateEventInput(params)
	if err != nil {
		return nil, err
	}
	// json parse of the json body
	bytes, err := params.Data.MarshalJSON()
	if err != nil {
		return nil, err
	}
	// conversion to datatypes.JSON
	jsonData := datatypes.JSON(bytes)
	// forming input params for service usge
	input := &dto.EventCreateParams{
		Email:       params.Email,
		Environment: params.Environment,
		Component:   params.Component,
		Message:     params.MessageString,
		Data:        jsonData,
	}

	// crete event by invoking event service
	event, err := e.service.CreateEvent(*input)
	if err != nil {
		return nil, err
	}
	// convert event struct to protobuf message
	return event.ConvertToMessage(), nil
}

func (e *eventServer) FilterEvents(ctx context.Context, input *pb.EventFilterInput) (*pb.EventFilterResponse, error) {
	// initialize the event filter
	eventFilter := filters.GetEventFilter(e.repo)
	log.Println(fmt.Sprintf("receiving request with params %v", input))
	// Building the filter params by following builder pattern
	// setting up different query params
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
	// get the data from db
	events := eventFilter.Get()
	// slice conversion from event array to EventResponse array
	eventsMap := make([]*pb.EventResponse, len(events))
	for i, event := range events {
		eventsMap[i] = event.ConvertToMessage()
	}
	// forming event filter response
	return &pb.EventFilterResponse{Events: eventsMap}, nil
}

// GetNewEventServer Stub for event server return EventServer intreface
func GetNewEventServer(db *gorm.DB) pb.EventServer {
	// initializing repo for DI (dependency injection)
	repo := repository.GetEventRepository(db)
	// init service
	eventService := services.GetEventService(repo)
	// return eventServer
	return &eventServer{
		repo:    repo,
		service: eventService,
	}
}