package controllers

import (
	"event-tracking/metrics"
	_ "event-tracking/metrics"
	"event-tracking/utils"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus custom registry
var (
	customRegistry = prometheus.NewRegistry()
	memoryStats    = new(runtime.MemStats)
)

func init() {
	// Register custom metrics with Prometheus
	prometheus.MustRegister(metrics.TotalHttpRequests)
	prometheus.MustRegister(metrics.TotalEventProcessed)
	prometheus.MustRegister(metrics.EventProcessingDuration)
	prometheus.MustRegister(metrics.Total2xxHttpRequests)
	prometheus.MustRegister(metrics.Total3xxHttpRequests)
	prometheus.MustRegister(metrics.Total4xxHttpRequests)
	prometheus.MustRegister(metrics.Total5xxHttpRequests)
	prometheus.MustRegister(metrics.MemoryUsageBytes)
	prometheus.MustRegister(metrics.MemoryUsagePercentage)
	prometheus.MustRegister(metrics.CPUUsagePercentage)
	prometheus.MustRegister(metrics.TotalGoroutines)
	prometheus.MustRegister(metrics.TotalThreads)

	// Register custom metrics with custom registry
	customRegistry.MustRegister(metrics.TotalHttpRequests)
	customRegistry.MustRegister(metrics.TotalEventProcessed)
	customRegistry.MustRegister(metrics.EventProcessingDuration)
	customRegistry.MustRegister(metrics.Total2xxHttpRequests)
	customRegistry.MustRegister(metrics.Total3xxHttpRequests)
	customRegistry.MustRegister(metrics.Total4xxHttpRequests)
	customRegistry.MustRegister(metrics.Total5xxHttpRequests)
	customRegistry.MustRegister(metrics.MemoryUsageBytes)
	customRegistry.MustRegister(metrics.MemoryUsagePercentage)
	customRegistry.MustRegister(metrics.CPUUsagePercentage)
	customRegistry.MustRegister(metrics.TotalGoroutines)
	customRegistry.MustRegister(metrics.TotalThreads)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		runtime.ReadMemStats(memoryStats)

		// Trigger HTTP requests total
		utils.TriggerHttpRequestsTotal(context.Request.Method, context.Request.URL.Path, strconv.Itoa(context.Writer.Status()))

		// Trigger 2XX HTTP requests total
		if context.Writer.Status() >= 200 && context.Writer.Status() < 300 {
			utils.Trigger2xxHttpRequestsTotal(context.Request.Method, context.Request.URL.Path)
		}

		// Trigger 3XX HTTP requests total
		if context.Writer.Status() >= 300 && context.Writer.Status() < 400 {
			utils.Trigger3xxHttpRequestsTotal(context.Request.Method, context.Request.URL.Path)
		}

		// Trigger 4XX HTTP requests total
		if context.Writer.Status() >= 400 && context.Writer.Status() < 500 {
			utils.Trigger4xxHttpRequestsTotal(context.Request.Method, context.Request.URL.Path)
		}

		// Trigger 5XX HTTP requests total
		if context.Writer.Status() >= 500 && context.Writer.Status() < 600 {
			utils.Trigger5xxHttpRequestsTotal(context.Request.Method, context.Request.URL.Path)
		}

		// Trigger memory usage in bytes
		utils.TriggerMemoryUsageBytes(float64(memoryStats.Alloc))

		// Trigger memory usage in percentage
		utils.TriggerMemoryUsagePercentage(float64(memoryStats.Alloc) / float64(memoryStats.Sys) * 100)

		// Trigger CPU usage in percentage
		utils.TriggerCPUUsagePercentage(float64(runtime.NumCPU()) / float64(runtime.NumCPU()) * 100)

		// Trigger total goroutines
		utils.TriggerTotalGoroutines(float64(runtime.NumGoroutine()))

		// Trigger total threads
		utils.TriggerTotalThreads(float64(runtime.NumCPU()))

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
