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

	// Total 2XX HTTP requests
	Total2xxHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_2xx_total",
			Help: "Total number of 2XX HTTP requests",
		},
		[]string{"method", "path"},
	)

	// Total 3XX HTTP requests
	Total3xxHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_3xx_total",
			Help: "Total number of 3XX HTTP requests",
		},
		[]string{"method", "path"},
	)

	// Total 4XX HTTP requests
	Total4xxHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_4xx_total",
			Help: "Total number of 4XX HTTP requests",
		},
		[]string{"method", "path"},
	)

	// Total 5XX HTTP requests
	Total5xxHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_5xx_total",
			Help: "Total number of 5XX HTTP requests",
		},
		[]string{"method", "path"},
	)

	// Memory usage in bytes
	MemoryUsageBytes = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Memory usage in bytes",
		},
	)

	// Memory usage in percentage
	MemoryUsagePercentage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_percentage",
			Help: "Memory usage in percentage",
		},
	)

	// CPU usage in percentage
	CPUUsagePercentage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percentage",
			Help: "CPU usage in percentage",
		},
	)

	// Total goroutines
	TotalGoroutines = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "total_goroutines",
			Help: "Total number of goroutines",
		},
	)

	// Total threads
	TotalThreads = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "total_threads",
			Help: "Total number of threads",
		},
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
