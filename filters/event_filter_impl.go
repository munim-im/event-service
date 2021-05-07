package filters

import (
	"event-service/interfaces/filters"
	"event-service/interfaces/repository"
	"event-service/models"
	"gorm.io/gorm"
	"strings"
)

type eventFilter struct {
	repo repository.EventRepository
	db *gorm.DB
}

func (e *eventFilter) Since(date *string) filters.EventFilter {
	if date != nil && *date != "" {
		e.db = e.db.Where("date_recorded >= ?", *date)
	}
	return e
}

func (e *eventFilter) Author(email *string) filters.EventFilter {
	if email != nil && *email != "" {
		e.db = e.db.Where("email = ?", *email)
	}
	return e
}

func (e *eventFilter) OfComponent(component *string) filters.EventFilter {
	if component != nil && *component != "" {
		e.db = e.db.Where("component = ?", *component)
	}
	return e
}

func (e *eventFilter) InEnvironment(environment *string) filters.EventFilter {
	if environment != nil && *environment != "" {
		e.db = e.db.Where("environment = ?", *environment)
	}
	return e
}

func (e *eventFilter) WithMessage(message *string) filters.EventFilter {
	if message != nil && *message != "" {
		e.db = e.db.Where("message ilike ?", strings.Join([]string {"%", *message, "%"}, ""))
	}
	return e
}

func (e *eventFilter) Get() []models.Event {
	var events []models.Event
	result := e.db.Find(&events)
	if result.Error != nil {
		return nil
	}
	return events
}

func GetEventFilter(eventRepository repository.EventRepository) filters.EventFilter {
	return &eventFilter{
		repo: eventRepository,
		db: eventRepository.GetDB(),
	}
}
