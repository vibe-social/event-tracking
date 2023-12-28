package controllers

import (
	"context"
	"net/http"
	"time"

	"event-tracking/database"
	"event-tracking/kafka"
	"event-tracking/models"
	"event-tracking/proto"
	"event-tracking/utils"

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
// @Param event body models.CreateEventRequest true "Event"
// @Router /events [post]
func CreateEvent(context *gin.Context) {
	// Start measuring the time
	timer := time.Now()

	var request models.CreateEventRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{Type: models.EventType(request.Type), UserId: request.UserId, Content: request.Content}
	database.DB.Create(&event)

	context.JSON(http.StatusOK, gin.H{"data": event})

	// Stop measuring the time and calculate the duration
	duration := time.Since(timer).Seconds()

	// Trigger the total event processed metric
	utils.TriggerTotalEventProcessed(models.EventType(event.Type))

	// Trigger the event processing duration metric
	utils.TriggerEventProcessingDuration(models.EventType(event.Type), duration)
}

// @Tags events
// @ID update-event
// @Summary Update event
// @Description update event
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Param event body models.UpdateEventRequest true "Event"
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
	kafkaRuntime, err := kafka.EH.GetRuntimeInformation(context)
	if err != nil {
		return nil, err
	}

	// Start measuring the time
	timer := time.Now()

	// Convert event to byte[]
	eventBytes := []byte(event.String())

	// Send the event to Azure Event Hub
	err = kafka.EH.Send(context, eventhub.NewEvent(eventBytes))

	// Stop measuring the time and calculate the duration
	duration := time.Since(timer).Seconds()

	// Trigger the total event processed metric
	utils.TriggerTotalEventProcessed(models.EventType(event.Type))

	// Trigger the event processing duration metric
	utils.TriggerEventProcessingDuration(models.EventType(event.Type), duration)

	if err != nil {
		// Trigger the Kafka outgoing errors metric
		utils.TriggerKafkaOutgoingErrors(kafkaRuntime.Path)
		return nil, err
	} else {
		// Trigger the Kafka outgoing requests metric
		utils.TriggerKafkaOutgoingRequests(kafkaRuntime.Path)

		// Trigger the Kafka outgoing bytes metric
		utils.TriggerKafkaOutgoingBytes(kafkaRuntime.Path, float64(len(eventBytes)))
	}

	return event, nil
}
