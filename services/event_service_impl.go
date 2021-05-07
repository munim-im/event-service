package services

import (
	"event-service/dto"
	"event-service/filters"
	"event-service/interfaces/repository"
	"event-service/interfaces/services"
	"event-service/models"
)

type eventService struct {
	eventRepo repository.EventRepository
}

func (e *eventService) FilterEvents(params dto.EventFilterParams) []models.Event {
	eventFilter := filters.GetEventFilter(e.eventRepo)
	eventFilter = eventFilter.Author(&params.Email).InEnvironment(&params.Environment).OfComponent(&params.Component).Since(&params.From).WithMessage(&params.Message)
	events := eventFilter.Get()
	return events
}

func (e *eventService) CreateEvent(params dto.EventCreateParams) (*models.Event, error) {
	event := &models.Event{
		Email:       params.Email,
		Environment: params.Environment,
		Component:   params.Component,
		Message:     params.Message,
		Data:        params.Data,
	}

	return e.eventRepo.CreateEvent(event)
}

func GetEventService(eventRepository repository.EventRepository) services.EventService {
	return &eventService{
		eventRepo: eventRepository,
	}
}
