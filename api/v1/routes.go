package v1

import (
	"event-service/api/v1/controllers"
	"event-service/repository"
	"event-service/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigureEventRoutesV1(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	v1 := r.Group("/v1/events")

	eventRepo := repository.GetEventRepository(db)
	eventService := services.GetEventService(eventRepo)
	v1Controller := controllers.GetEventControllerV1(eventService)

	// route mapping with url
	v1.POST("", v1Controller.CreateEvent)
	v1.GET("", v1Controller.RetrieveEvent)

	return r
}