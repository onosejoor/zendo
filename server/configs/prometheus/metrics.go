package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ProjectsCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_zendo_projects_created",
		Help: "Total number of projects created",
	})

	TasksCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_zendo_tasks_created",
		Help: "Total number of tasks created",
	})

	UsersRegistered = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_zendo_users_signed_up",
		Help: "Total number of users who signed up",
	})

	DatabaseQueryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "zendo_database_query_duration_seconds",
		Help:    "Duratio of database queries",
		Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 1, 2},
	}, []string{"collection", "operation"})

	RedisOperations = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "total_zendo_redis_operations",
		Help: "Total number of Redis operations",
	}, []string{"operation"})

	CronJobsOperation = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "total_zendo_cron_jobs_operations",
		Help: "Total number of Cron Jobs operations",
	}, []string{"operation"})
)

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
