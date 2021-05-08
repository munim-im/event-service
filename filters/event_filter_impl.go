package filters

import (
	"event-service/interfaces/filters"
	"event-service/interfaces/repository"
	"event-service/models"
	"gorm.io/gorm"
	"strings"
)
// eventFilter - Event Filter instance
type eventFilter struct {
	// interface definition
	filters.EventFilter
	// repo instance
	repo repository.EventRepository
	// db instance
	db *gorm.DB
}

// Since override
func (e *eventFilter) Since(date *string) filters.EventFilter {
	// dynamic query formation of date
	if date != nil && *date != "" {
		e.db = e.db.Where("date_recorded >= ?", *date)
	}
	return e
}

// Author override
func (e *eventFilter) Author(email *string) filters.EventFilter {
	// dynamic query formation of email
	if email != nil && *email != "" {
		e.db = e.db.Where("email = ?", *email)
	}
	return e
}

// OfComponent override
func (e *eventFilter) OfComponent(component *string) filters.EventFilter {
	// dynamic query formation of component
	if component != nil && *component != "" {
		e.db = e.db.Where("component = ?", *component)
	}
	return e
}

// InEnvironment override
func (e *eventFilter) InEnvironment(environment *string) filters.EventFilter {
	// dynamic query formation of environment
	if environment != nil && *environment != "" {
		e.db = e.db.Where("environment = ?", *environment)
	}
	return e
}

// WithMessage override
func (e *eventFilter) WithMessage(message *string) filters.EventFilter {
	// dynamic query formation of message for partial filtering
	if message != nil && *message != "" {
		e.db = e.db.Where("message ilike ?", strings.Join([]string {"%", *message, "%"}, ""))
	}
	return e
}

// Get override
func (e *eventFilter) Get() []models.Event {
	// holder for event list
	var events []models.Event
	result := e.db.Find(&events)
	if result.Error != nil {
		return nil
	}
	return events
}

// GetEventFilter functions for getting event filter interface
func GetEventFilter(eventRepository repository.EventRepository) filters.EventFilter {
	return &eventFilter{
		repo: eventRepository,
		db: eventRepository.GetDB(),
	}
}
