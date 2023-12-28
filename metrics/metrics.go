package metrics

import "github.com/prometheus/client_golang/prometheus"

// Define Prometheus custom metrics
var (
	// Total HTTP requests
	TotalHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// Total Event Processed
	TotalEventProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_event_processed",
			Help: "Total number of events processed",
		},
		[]string{"event_type"},
	)

	// Event Processing Duration
	EventProcessingDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "event_processing_duration",
			Help:    "Event processing duration",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"event_type"},
	)
)
