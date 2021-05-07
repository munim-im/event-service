package filters

import (
	"event-service/models"
)

type EventFilter interface {
	Get() []models.Event
	Since(date *string) EventFilter
	Author(email *string) EventFilter
	OfComponent(component *string) EventFilter
	InEnvironment(environment *string) EventFilter
	WithMessage(message *string) EventFilter
}
