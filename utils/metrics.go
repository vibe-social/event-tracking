package utils

import (
	"event-tracking/metrics"
	"event-tracking/models"
)

func TriggerHttpRequestsTotal(method string, path string, status string) {
	// Increment HTTP requests total
	metrics.TotalHttpRequests.WithLabelValues(method, path, status).Inc()
}

func TriggerTotalEventProcessed(eventType models.EventType) {
	// Increment total event processed
	metrics.TotalEventProcessed.WithLabelValues(string(eventType)).Inc()
}

func TriggerEventProcessingDuration(eventType models.EventType, duration float64) {
	// Observe event processing duration
	metrics.EventProcessingDuration.WithLabelValues(string(eventType)).Observe(duration)
}
