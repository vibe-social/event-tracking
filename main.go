package main

import (
	controllers "event-tracking/controllers"
	"event-tracking/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Specify the routes and the controllers
	router.GET("/events", controllers.FindEvents)
	router.GET("/events/:id", controllers.FindEvent)
	router.POST("/events", controllers.CreateEvent)
	router.PATCH("/events/:id", controllers.UpdateEvent)
	router.DELETE("/events/:id", controllers.DeleteEvent)

	// Run the server
	router.Run()
}
