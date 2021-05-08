package db

// interface for db connection
type DBConnection interface {
	GetDSNString() *string
}
