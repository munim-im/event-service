package dto

import "gorm.io/datatypes"

type EventCreateParams struct {
	Email       string         `json:"email" binding:"required,email"`
	Environment string         `json:"environment" binding:"required"`
	Component   string         `json:"component" binding:"required"`
	Message     string         `json:"message" binding:"required"`
	Data        datatypes.JSON `json:"data"`
}
