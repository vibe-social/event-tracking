package controllers

import (
	"event-tracking/metrics"
	_ "event-tracking/metrics"
	"event-tracking/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus custom registry
var (
	customRegistry = prometheus.NewRegistry()
)

func init() {
	// Register custom metrics with Prometheus
	prometheus.MustRegister(metrics.TotalHttpRequests)
	prometheus.MustRegister(metrics.TotalEventProcessed)
	prometheus.MustRegister(metrics.EventProcessingDuration)

	// Register custom metrics with custom registry
	customRegistry.MustRegister(metrics.TotalHttpRequests)
	customRegistry.MustRegister(metrics.TotalEventProcessed)
	customRegistry.MustRegister(metrics.EventProcessingDuration)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Trigger HTTP requests total
		utils.TriggerHttpRequestsTotal(context.Request.Method, context.Request.URL.Path, strconv.Itoa(context.Writer.Status()))

		// Continue with the next middleware or route handler
		context.Next()
	}
}

// @Tags metrics
// @ID prometheus-metrics
// @Summary Prometheus metrics
// @Description Prometheus metrics
// @Produce  json
// @Router /metrics [get]
func PrometheusHandler() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(context *gin.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	}
}

// @Tags metrics
// @ID custom-prometheus-metrics
// @Summary Custom Prometheus metrics
// @Description Custom Prometheus metrics
// @Produce  json
// @Router /custom-metrics [get]
func CustomPrometheusHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{}).ServeHTTP(context.Writer, context.Request)
	}
}
