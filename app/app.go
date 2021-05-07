package app

import (
	"event-service/db"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	// explanation why the injection working
	// configure db connection in here
	dbConfig := db.GetPostgresConnection()
	dbInstance, _ := dbConfig.DB()
	defer dbInstance.Close() // graceful connection termination

	// initialize the routers
	r := gin.Default()
	r = ConfigureAppRoutes(r, dbConfig)
	r.Run()
}