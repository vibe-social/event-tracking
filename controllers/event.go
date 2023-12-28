package controllers

import (
	"context"
	"net/http"

	"event-tracking/database"
	"event-tracking/kafka"
	"event-tracking/models"
	"event-tracking/proto"

	eventhub "github.com/Azure/azure-event-hubs-go"
	"github.com/gin-gonic/gin"
)

// @Tags events
// @ID get-events
// @Summary List events
// @Description get events
// @Produce  json
// @Router /events [get]
func FindEvents(context *gin.Context) {
	var events []models.Event
	database.DB.Find(&events)

	context.JSON(http.StatusOK, gin.H{"data": events})
}

// @Tags events
// @ID get-event
// @Summary Get event
// @Description get event
// @Produce  json
// @Param id path int true "Event ID"
// @Router /events/{id} [get]
func FindEvent(context *gin.Context) {
	var event models.Event
	if err := database.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": event})
}

// @Tags events
// @ID create-event
// @Summary Create event
// @Description create event
// @Accept  json
// @Produce  json
// @Param event body CreateEventRequest true "Event"
// @Router /events [post]
func CreateEvent(context *gin.Context) {
	var request models.CreateEventRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{Type: request.Type, UserId: request.UserId, Content: request.Content}
	database.DB.Create(&event)

	context.JSON(http.StatusOK, gin.H{"data": event})
}

// @Tags events
// @ID update-event
// @Summary Update event
// @Description update event
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Param event body UpdateEventRequest true "Event"
// @Router /events/{id} [patch]
func UpdateEvent(context *gin.Context) {
	var event models.Event
	if err := database.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var request models.UpdateEventRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&event).Updates(request)

	context.JSON(http.StatusOK, gin.H{"data": event})
}

// @Tags events
// @ID delete-event
// @Summary Delete event
// @Description delete event
// @Produce  json
// @Param id path int true "Event ID"
// @Router /events/{id} [delete]
func DeleteEvent(context *gin.Context) {
	// Get model if exist
	var event models.Event
	if err := database.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&event)

	context.JSON(http.StatusOK, gin.H{"data": true})
}

type Server struct{}

// @Tags events
// @ID create-event-grpc
// @Summary Create event
// @Description create event
func (s *Server) CreateEvent(context context.Context, event *proto.Event) (*proto.Event, error) {
	// Convert event to byte[]
	eventBytes := []byte(event.String())

	// Send the event to Azure Event Hub
	kafka.EH.Send(context, eventhub.NewEvent(eventBytes))

	return event, nil
}
