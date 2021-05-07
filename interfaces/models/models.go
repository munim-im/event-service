package models

type IdGenerator interface {
	GenerateId() *string
}
