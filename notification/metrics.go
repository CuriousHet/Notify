package notification

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	CustomRegistry = prometheus.NewRegistry()

	NotificationsSent = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "notifications_sent_total",
			Help: "Total number of notifications successfully sent",
		},
	)
	NotificationsFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "notifications_failed_total",
			Help: "Total number of notifications that failed to send",
		},
	)
	DeliveryDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "notification_delivery_duration_seconds",
			Help:    "Duration of notification deliveries",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func InitMetrics() {
	CustomRegistry.MustRegister(NotificationsSent)
	CustomRegistry.MustRegister(NotificationsFailed)
	CustomRegistry.MustRegister(DeliveryDuration)
}
