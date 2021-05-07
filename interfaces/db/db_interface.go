package db

type DBConnection interface {
	GetDSNString() *string
}
