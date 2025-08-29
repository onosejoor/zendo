package prometheus

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var initOnce sync.Once

var (
	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 3, 5, 10},
	}, []string{"path", "method", "status"})

	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests in total",
	}, []string{"path", "method", "status"})

	HttpRequestInProgress = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_requests_in_progress",
		Help: "Total number of HTTP requests in progress",
	}, []string{"path", "method"})
)

var (
	ProjectsCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_zendo_projects_created",
		Help: "Total number of projects created",
	})

	TasksCreated = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_zendo_tasks_created",
		Help: "Total number of tasks created",
	})

	UsersRegistered = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_zendo_users_signed_up",
		Help: "Total number of users who signed up",
	})

	DatabaseQueryDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "zendo_database_query_duration_seconds",
		Help:    "Duratio of database queries",
		Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 1, 2},
	}, []string{"collection", "operation"})

	RedisOperations = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "total_zendo_redis_operations",
		Help: "Total number of Redis operations",
	}, []string{"operation"})

	CronJobsOperation = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "total_zendo_cron_jobs_operations",
		Help: "Total number of Cron Jobs operations",
	}, []string{"operation"})
)

func Init() {
	initOnce.Do(func() {
		prometheus.MustRegister(httpRequestDuration)
		prometheus.MustRegister(HttpRequestTotal)
		prometheus.MustRegister(HttpRequestInProgress)
		prometheus.MustRegister(ProjectsCreated)
		prometheus.MustRegister(TasksCreated)
		prometheus.MustRegister(UsersRegistered)
		prometheus.MustRegister(DatabaseQueryDuration)
		prometheus.MustRegister(RedisOperations)
		prometheus.MustRegister(CronJobsOperation)
	})
}

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

func RecordProjectCreation() {
	ProjectsCreated.Inc()
}

func RecordTaskCreation() {
	TasksCreated.Inc()
}

func RecordUserRegistration() {
	UsersRegistered.Inc()
}

func RecordDatabaseQueryOperation(duration time.Duration, collection, operation string) {
	DatabaseQueryDuration.WithLabelValues(collection, operation).Observe(duration.Seconds())
}

func RecordRedisOperation(operation string) {
	RedisOperations.WithLabelValues(operation).Inc()
}

func RecordCronJobOperation(operation string) {
	CronJobsOperation.WithLabelValues(operation).Inc()
}
