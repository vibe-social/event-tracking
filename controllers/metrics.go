package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// Register custom metrics with Prometheus
	prometheus.MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	}))
}

// @Tags metrics
// @ID prometheus-metrics
// @Summary Prometheus metrics
// @Description Prometheus metrics
// @Produce  json
// @Router /metrics [get]
func PrometheusHandler() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
