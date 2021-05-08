package models

// interface for id generation
type IdGenerator interface {
	GenerateId() *string
}
