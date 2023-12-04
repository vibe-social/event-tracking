package main

import (
	controllers "event-tracking/controllers"
	"event-tracking/models"

	_ "event-tracking/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

func init() {
	prometheus.MustRegister(cpuTemp)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// @title Event Tracking API Documentation
// @description Event tracking documentation for the social networking app focused on sharing vibes.
// @version 1.0
// @contact.name Rok Mokotar
// @contact.url https://www.linkedin.com/in/mokot/
// @contact.email rm6551@student.uni-lj.si
func main() {
	// Create a default gin router
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Swagger documentation endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/metrics", prometheusHandler())

	// Health check endpoint
	router.GET("/health", controllers.CheckHealth)

	// Specify the events routes and the controllers
	router.GET("/events", controllers.FindEvents)
	router.GET("/events/:id", controllers.FindEvent)
	router.POST("/events", controllers.CreateEvent)
	router.PATCH("/events/:id", controllers.UpdateEvent)
	router.DELETE("/events/:id", controllers.DeleteEvent)

	// TODO - Add a route to handle metrics

	// Run the server
	router.Run(":8080")
}
