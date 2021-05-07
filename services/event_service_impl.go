package services

import (
	"event-service/dto"
	"event-service/interfaces/repository"
	"event-service/interfaces/services"
	"event-service/models"
)

type eventService struct {
	eventRepo repository.EventRepository
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
