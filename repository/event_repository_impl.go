package repository

import (
	"event-service/interfaces/repository"
	"event-service/models"
	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func (e *eventRepository) GetDB() *gorm.DB {
	return e.db
}

func (e *eventRepository) GetAllEvent() ([]models.Event, error) {
	var events []models.Event
	result := e.db.Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func (e *eventRepository) CreateEvent(event *models.Event) (*models.Event, error) {
	result := e.db.Create(&event)
	if result.Error != nil {
		return nil, result.Error
	}
	return event, nil
}

func GetEventRepository(db *gorm.DB) repository.EventRepository{
	return &eventRepository{
		db: db,
	}
}
