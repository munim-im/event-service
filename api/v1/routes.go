package v1

import (
	"event-service/api/v1/controllers"
	"event-service/repository"
	"event-service/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigureEventRoutesV1(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	// defining the route group
	v1 := r.Group("/v1/events")

	/*
	Initializing repo, service and controller interface so that our routes have everything setup in one place
	 */

	// initialize repository instance with db from dependency
	eventRepo := repository.GetEventRepository(db)
	// inject repo as depenency to the service
	eventService := services.GetEventService(eventRepo)
	// inject service to the controllers
	v1Controller := controllers.GetEventControllerV1(eventService)

	// route mapping with url
	// Post request for creating event
	v1.POST("", v1Controller.CreateEvent)
	// Get requests for filtering
	v1.GET("", v1Controller.RetrieveEvent)

	return r
}