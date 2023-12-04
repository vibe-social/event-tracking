package controllers

import (
	"net/http"

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

// @Summary List events
// @Description get events
// @ID get-events
// @Accept  json
// @Produce  json
// @Router /events [get]
func FindEvents(context *gin.Context) {
	var events []models.Event
	models.DB.Find(&events)

	context.JSON(http.StatusOK, gin.H{"data": events})
}

func FindEvent(context *gin.Context) {
	var event models.Event
	if err := models.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func CreateEvent(context *gin.Context) {
	var request CreateEventRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{Type: request.Type, UserId: request.UserId, Content: request.Content}
	models.DB.Create(&event)

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	if err := models.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var request UpdateEventRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&event).Updates(request)

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func DeleteEvent(context *gin.Context) {
	// Get model if exist
	var event models.Event
	if err := models.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&event)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
