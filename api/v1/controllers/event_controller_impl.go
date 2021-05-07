package controllers

import (
	"event-service/dto"
	"event-service/interfaces/controllers"
	"event-service/interfaces/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type eventController struct {
	eventService services.EventService
}

func (e *eventController) CreateEvent(c *gin.Context) {
	var body dto.EventCreateParams
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	event, err := e.eventService.CreateEvent(body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (e *eventController) RetrieveEvent(c *gin.Context) {
	var query dto.EventFilterParams
	err := c.ShouldBindQuery(&query)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	events := e.eventService.FilterEvents(query)
	c.JSON(http.StatusOK, events)
}

func GetEventControllerV1(service services.EventService) controllers.EventController{
	return &eventController{
		eventService: service,
	}
}
