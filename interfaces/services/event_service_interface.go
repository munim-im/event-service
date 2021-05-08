package services

import (
	"event-service/dto"
	"event-service/models"
)
// interface for event service
type EventService interface {
	CreateEvent(params dto.EventCreateParams) (*models.Event, error)
	FilterEvents(params dto.EventFilterParams) []models.Event
}
