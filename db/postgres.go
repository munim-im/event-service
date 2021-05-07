package db

import (
	"event-service/dto"
	"event-service/models"
	"event-service/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresConfig struct {
	dto.DBConfig
}

func (r *PostgresConfig) GetDSNString() *string {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		r.Host, r.Username, r.Password, r.Name, r.Port, r.Timezone,
		)
	return &dsn
}

func getDBConfig() *PostgresConfig {
	return &PostgresConfig{struct {
		Username string
		Password string
		Host     string
		Port     string
		Name     string
		Timezone string
	}{
		Username: utils.GetEnvironmentVariable("DB_USER", "nashmaniac"),
		Password: utils.GetEnvironmentVariable("DB_PASSWORD", "shetu2153"),
		Host: utils.GetEnvironmentVariable("DB_HOST", "localhost"),
		Port: utils.GetEnvironmentVariable("DB_PORT", "5432"),
		Name: utils.GetEnvironmentVariable("DB_NAME", "event-service"),
		Timezone: utils.GetEnvironmentVariable("DB_TIMEZONE", "Asia/Dhaka"),
	}}
}


func GetPostgresConnection() *gorm.DB {
	config := getDBConfig()
	dsnString := *config.GetDSNString()
	db, err := gorm.Open(postgres.Open(dsnString), &gorm.Config{})
	if err != nil {
		panic("Error in connecting the database")
	}
	log.Println(fmt.Sprintf("Connection to %v is successful", dsnString))
	// we can configure database resolvers here to use for more database if required
	// run the migrations in here
	// right now using the gorm default one. For scaling would use gormigrate or golang-migrate/migrate packages
	db.AutoMigrate(&models.Event{})
	return db
}
