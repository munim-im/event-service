package models

import (
	"event-service/config"
	"event-service/interfaces/models"
	"event-service/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Event struct {
	models.IdGenerator `gorm:"-" json:"-"`
	ID                 string         `json:"id" gorm:"primaryKey;index;unique;not null;type:varchar(50)"`
	Email              string         `json:"email" gorm:"type:varchar(100);not null"`
	Environment        string         `json:"environment" gorm:"not null;type:varchar(200)"`
	Component          string         `json:"component" gorm:"not null;type:varchar(300)"`
	Message            string         `json:"message" gorm:"not null"`
	Data               datatypes.JSON `json:"data" gorm:"DEFAULT:null;"`
	DateRecorded       datatypes.Date `json:"date_recorded" gorm:"not null"`
	CreatedAt          time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt          time.Time      `json:"updated_at" gorm:"not null"`
}

// using hooks to setup the created time and edited time
func (e *Event) GenerateId() *string {
	environment := strings.ToLower(e.Environment)[:3]
	component := strings.ToLower(e.Component)[:3]
	currentTime := time.Now()
	timestamp := currentTime.Format(config.TIMESTAMP_FORMAT_ID)
	id := strings.Join([]string{environment, component, timestamp, utils.GenerateHash(currentTime.Nanosecond(), 3)}, "-")
	return &id
}

func (u *Event) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = *u.GenerateId()

	// handle empty json data
	if u.Data == nil {
		u.Data = datatypes.JSON([]byte(`{}`))
	}
	u.DateRecorded = datatypes.Date(time.Now())
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return
}

func (u *Event) BeforeSave(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
