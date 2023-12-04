package main

import (
	controllers "event-tracking/controllers"
	"event-tracking/models"

	_ "event-tracking/docs"

	"github.com/gin-gonic/gin"
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
	// Create a default gin router
	router := gin.Default()

	// Add the swagger endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Connect to database
	models.ConnectDatabase()

	// Specify the routes and the controllers
	router.GET("/events", controllers.FindEvents)
	router.GET("/events/:id", controllers.FindEvent)
	router.POST("/events", controllers.CreateEvent)
	router.PATCH("/events/:id", controllers.UpdateEvent)
	router.DELETE("/events/:id", controllers.DeleteEvent)

	// TODO - Add a route to handle metrics

	// Run the server
	router.Run(":8080")
}
