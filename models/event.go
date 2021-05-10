package models

import (
	"encoding/json"
	"event-service/config"
	"event-service/interfaces/models"
	"event-service/pb"
	"event-service/utils"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
	"time"
)
// Event - database model for event
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

// id generator override
func (e *Event) GenerateId() *string {
	environment := strings.ToLower(e.Environment)[:3]
	component := strings.ToLower(e.Component)[:3]
	currentTime := time.Now()
	timestamp := currentTime.Format(config.TIMESTAMP_FORMAT_ID)
	id := strings.Join([]string{environment, component, timestamp, utils.GenerateHash(currentTime.Nanosecond(), 3)}, "-")
	return &id
}
// using hooks to setup the created time and edited time
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
// before save hook
func (u *Event) BeforeSave(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (e Event) ConvertToMessage() *pb.EventResponse {
	bytes, _ := e.Data.MarshalJSON()
	var p map[string]interface{};
	_ = json.Unmarshal(bytes, &p)
	jsonValue, _ := structpb.NewStruct(p)

	return &pb.EventResponse{
		Id:            e.ID,
		Email:         e.Email,
		Environment:   e.Environment,
		Component:     e.Component,
		MessageString: e.Message,
		Data:          jsonValue,
	}
}
