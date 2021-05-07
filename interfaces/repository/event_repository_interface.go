package repository

import (
	"event-service/models"
)

type EventRepository interface {
	CreateEvent(event *models.Event) (*models.Event, error)
}
