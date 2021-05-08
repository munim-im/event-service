package app

import (
	"event-service/db"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	// configure db connection in here
	dbConfig := db.GetPostgresConnection()
	dbInstance, _ := dbConfig.DB()
	defer dbInstance.Close() // graceful connection termination

	// initialize the routers
	r := gin.Default()
	// configure application routes
	r = ConfigureAppRoutes(r, dbConfig)
	// run the servers
	r.Run()
}