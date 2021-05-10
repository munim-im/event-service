package app

import (
	"event-service/db"
	"event-service/pb"
	"event-service/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func setupGinServer() {
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

func setUpGRPCServer() {
	dbConfig := db.GetPostgresConnection()
	dbInstance, _ := dbConfig.DB()
	defer dbInstance.Close()
	grpcServer := grpc.NewServer()
	eventServer := server.GetNewEventServer(dbConfig)
	pb.RegisterEventServer(grpcServer, eventServer)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err.Error())
	}
	log.Println(fmt.Sprintf("Serving at location %v", listener.Addr().String()))
	grpcServer.Serve(listener)
}

func StartApp() {
	//setupGinServer()
	setUpGRPCServer()
}