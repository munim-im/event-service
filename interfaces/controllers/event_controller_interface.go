package controllers

import "github.com/gin-gonic/gin"

// explain why we keep it centralized

type EventController interface {
	CreateEvent(c *gin.Context)
	RetrieveEvent(c *gin.Context)
}