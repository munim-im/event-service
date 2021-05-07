package services

import (
	"event-service/dto"
	"event-service/models"
)

type EventService interface {
	CreateEvent(params dto.EventCreateParams) (*models.Event, error)
}
