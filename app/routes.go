package app

import (
	v1 "event-service/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func configureEventRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r = v1.ConfigureEventRoutesV1(r, db)
	return r
}

func ConfigureAppRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	apiGroup := r.Group("/api")
	apiGroup = configureEventRoutes(apiGroup, db)
	return r
}