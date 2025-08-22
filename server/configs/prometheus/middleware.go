package middlewares

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 3, 5, 10},
	}, []string{"path", "method", "status"})

	HttpRequestTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests in progress",
	}, []string{"path", "method"})

	HttpRequestInProgress = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_requests_in_progress",
		Help: "Number of HTTP requests in progress",
	}, []string{"path", "method"})
)

func NewMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		path := c.Route().Path
		method := c.Method()

		HttpRequestInProgress.WithLabelValues(path, method).Inc()
		err := c.Next()

		HttpRequestInProgress.WithLabelValues(path, method).Dec()
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Response().StatusCode())
		httpRequestDuration.WithLabelValues(path, method, status).Observe(duration)
		HttpRequestTotal.WithLabelValues(path, method, status).Inc()
		return err
	}
}
