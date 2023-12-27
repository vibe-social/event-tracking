package main

import (
	"event-tracking/configs"
	"event-tracking/controllers"
	"event-tracking/database"
	"fmt"
	"log"

	_ "event-tracking/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Set the router mode
	routerMode := viper.GetString("SERVER_MODE")
	if routerMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if routerMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create a default gin router
	router := gin.Default()

	// Connect to database
	database.ConnectDatabase()

	// Swagger documentation endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Prometheus metrics endpoint
	router.GET("/metrics", controllers.PrometheusHandler())

	// Health check endpoint
	router.GET("/health", controllers.CheckHealth)

	// Specify the events routes and the controllers
	router.GET("/events", controllers.FindEvents)
	router.GET("/events/:id", controllers.FindEvent)
	router.POST("/events", controllers.CreateEvent)
	router.PATCH("/events/:id", controllers.UpdateEvent)
	router.DELETE("/events/:id", controllers.DeleteEvent)

	// Run the server
	address := fmt.Sprintf(":%d", viper.GetInt("SERVER_PORT"))
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
