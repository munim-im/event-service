package repository

import (
	"event-service/models"
	"gorm.io/gorm"
)

// repository interface for general query
type EventRepository interface {
	GetDB() *gorm.DB
	CreateEvent(event *models.Event) (*models.Event, error)
	GetAllEvent() ([]models.Event, error)
}
