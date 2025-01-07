package metrics

import (
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type OrderMetrics struct {
	logger          slog.Logger
	service         string
	ReqByStatusCode *prometheus.CounterVec
	Duration        *prometheus.HistogramVec
}

func NewOrderMetrics(l slog.Logger, reg prometheus.Registerer) *OrderMetrics {
	service := "of-order-service"
	m := &OrderMetrics{
		logger:  l,
		service: service,
		ReqByStatusCode: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "requests_status_code",
			Help: "Requests by status code"},
			[]string{"service", "status"}),
		Duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Duration of request",
			Buckets: []float64{0.0, 0.001, 0.002, 0.003, 0.005, 0.007, 0.009, 0.01, 0.015, 0.02, 0.023, 0.025, 0.027, 0.029, 0.03, 0.031, 0.033, 0.035, 0.04, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3}},
			[]string{"service", "status", "method", "uri"}),
	}

	reg.MustRegister(m.ReqByStatusCode, m.Duration)
	return m
}

func (m *OrderMetrics) IncReqByStatusCode(status string) {
	m.ReqByStatusCode.With(prometheus.Labels{"service": m.service, "status": status}).Inc()
}

func (m *OrderMetrics) MeasureDuration(start time.Time, method string, uri string, statusCode string) {
	m.Duration.With(prometheus.Labels{
		"service": m.service,
		"method":  method,
		"uri":     uri,
		"status":  statusCode,
	}).Observe(float64(time.Since(start).Seconds()))
}
