package controllers

import (
	"event-service/dto"
	"event-service/interfaces/controllers"
	"event-service/interfaces/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type eventController struct {
	// interface implementation
	controllers.EventController
	// service object as dependency
	eventService services.EventService
}

// post request handler for creating events

func (e *eventController) CreateEvent(c *gin.Context) {
	var body dto.EventCreateParams
	// validating incoming body
	err := c.ShouldBindJSON(&body)
	// handling invalid input data
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	// creating event by invoking service
	event, err := e.eventService.CreateEvent(body)
	// handling errors from service
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	// returning the final event that is created
	c.JSON(http.StatusOK, event)
}

func (e *eventController) RetrieveEvent(c *gin.Context) {
	var query dto.EventFilterParams
	// binding the query params for filtering
	err := c.ShouldBindQuery(&query)
	// handling error during data binding
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	// filtering events by invoking service call
	events := e.eventService.FilterEvents(query)
	// returning data
	c.JSON(http.StatusOK, events)
}

// function to retrieve the controller interface

func GetEventControllerV1(service services.EventService) controllers.EventController{
	return &eventController{
		eventService: service,
	}
}
