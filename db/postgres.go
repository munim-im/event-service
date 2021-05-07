package db

import (
	"event-service/dto"
	"event-service/models"
	"event-service/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

// For generating dynamic hardcoded DSN string for multiple database
/***
func getDBConfigDynamicName(dbName string) *PostgresConfig  {
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
		Name: utils.GetEnvironmentVariable("DB_NAME", dbName),
		Timezone: utils.GetEnvironmentVariable("DB_TIMEZONE", "Asia/Dhaka"),
	}}
}
 */

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
	db, err := gorm.Open(postgres.Open(dsnString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Error in connecting the database")
	}
	log.Println(fmt.Sprintf("Connection to %v is successful", dsnString))

	// we can configure database resolvers here to use for more database if required
	// Source & Replica setup for read/write splitting

	/***

	source1 := getDBConfigDynamicName("event-service-source-1")
	replicas1 := getDBConfigDynamicName("event-service-replica-1")
	replicas2 := getDBConfigDynamicName("event-service-replica-2")
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{postgres.Open(*source1.GetDSNString()), postgres.Open(dsnString)},
		Replicas: []gorm.Dialector{postgres.Open(*replicas1.GetDSNString()), postgres.Open(*replicas2.GetDSNString())},
		Policy:   dbresolver.RandomPolicy{},
	}))

	 */

	// run the migrations in here
	// right now using the gorm default one. For scaling would use gormigrate or golang-migrate/migrate packages

	db.AutoMigrate(&models.Event{})
	return db
}
