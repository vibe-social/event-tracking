package controllers

import (
	"net/http"

	"event-tracking/database"
	"event-tracking/models"

	"github.com/gin-gonic/gin"
)

type CreateEventRequest struct {
	Type    string `json:"type" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateEventRequest struct {
	Type    string `json:"type"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}

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
	var request CreateEventRequest
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

	var request UpdateEventRequest
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
