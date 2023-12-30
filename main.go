package main

import (
	"event-tracking/configs"
	"event-tracking/controllers"
	"event-tracking/database"
	"event-tracking/kafka"
	"event-tracking/middleware"
	"event-tracking/proto"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	_ "event-tracking/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @title Event Tracking API Documentation
// @description Event tracking documentation for the social networking app focused on sharing vibes.
// @version 1.0
// @contact.name Rok Mokotar
// @contact.url https://www.linkedin.com/in/mokot/
// @contact.email rm6551@student.uni-lj.si
func main() {
	// Set configuration parameters
	configs.LoadConfig()

	// Set the server mode
	serverMode := viper.GetString("EVENT_TRACKING_SERVER_MODE")
	if serverMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if serverMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create the HTTP server
	httpRouter := gin.Default()

	// Create the gRPC server
	server := controllers.Server{}
	grpcServer := grpc.NewServer()

	// Connect to database
	database.ConnectDatabase()

	// Connect to Azure Event Hub
	kafka.ConnectEventHub()

	// Apply the Prometheus middleware
	httpRouter.Use(middleware.PrometheusMiddleware())

	// Swagger documentation endpoint
	httpRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	httpRouter.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Prometheus metrics endpoints
	httpRouter.GET("/metrics", controllers.PrometheusHandler())
	httpRouter.GET("/custom-metrics", controllers.CustomPrometheusHandler())

	// Specify the HTTP health endpoints and the controllers
	httpRouter.GET("/health", controllers.CheckHealth)
	httpRouter.GET("/health/general", controllers.CheckHealthGeneral)
	httpRouter.GET("/health/disk", controllers.CheckHealthDisk)
	httpRouter.GET("/health/cpu", controllers.CheckHealthCPU)
	httpRouter.GET("/health/goroutine", controllers.CheckHealthGoroutine)
	httpRouter.GET("/health/database", controllers.CheckHealthDatabase)
	httpRouter.GET("/health/kafka", controllers.CheckHealthKafka)
	httpRouter.GET("/health/live", controllers.CheckHealthLiveness)
	httpRouter.GET("/health/ready", controllers.CheckHealthReadiness)

	// Specify the HTTP events endpoints and the controllers
	httpRouter.GET("/events", controllers.FindEvents)
	httpRouter.GET("/events/:id", controllers.FindEvent)
	httpRouter.POST("/events", controllers.CreateEvent)
	httpRouter.PATCH("/events/:id", controllers.UpdateEvent)
	httpRouter.DELETE("/events/:id", controllers.DeleteEvent)

	// Specify the gRPC events endpoints and the controllers
	proto.RegisterEventServiceServer(grpcServer, &server)

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	// Create a WaitGroup and add a count of two, one for each goroutine
	var waitGroup sync.WaitGroup

	// Run the HTTP server in a separate goroutine
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		httpAddress := fmt.Sprintf(":%d", viper.GetInt("EVENT_TRACKING_HTTP_SERVER_PORT"))
		httpServer := &http.Server{
			Addr:         httpAddress,
			Handler:      httpRouter,
			ReadTimeout:  viper.GetDuration("EVENT_TRACKING_HTTP_SERVER_TIMEOUT"),
			WriteTimeout: viper.GetDuration("EVENT_TRACKING_HTTP_SERVER_TIMEOUT"),
		}
		log.Printf("HTTP server listening on port %s", httpAddress)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Run the gRPC server in a separate goroutine
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		grpcAddress := fmt.Sprintf(":%d", viper.GetInt("EVENT_TRACKING_GRPC_SERVER_PORT"))
		log.Printf("gRPC server listening on port %s", grpcAddress)
		grpcListener, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			log.Fatal(err)
		}
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for all goroutines to complete
	waitGroup.Wait()
}
