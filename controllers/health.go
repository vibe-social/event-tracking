package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags health
// @ID health-check
// @Summary Health check
// @Description health check
// @Produce  json
// @Router /health [get]
func CheckHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}
