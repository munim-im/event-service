package controllers

import "github.com/gin-gonic/gin"

// interface for event controller
type EventController interface {
	CreateEvent(c *gin.Context)
	RetrieveEvent(c *gin.Context)
}