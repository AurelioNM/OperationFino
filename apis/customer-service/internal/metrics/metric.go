package metrics

import (
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type CustomerMetrics struct {
	logger          slog.Logger
	service         string
	ReqByStatusCode *prometheus.CounterVec
	Duration        *prometheus.HistogramVec
}

func NewCustomerMetrics(l slog.Logger, reg prometheus.Registerer) *CustomerMetrics {
	service := "of-customer-service"
	m := &CustomerMetrics{
		logger:  l,
		service: service,
		ReqByStatusCode: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "requests_status_code",
			Help: "Requests by status code"},
			[]string{"service", "status"}),
		Duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Duration of request",
			Buckets: []float64{0.03, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.5, 0.6, 0.7, 0.8, 1.0, 1.5, 2.0, 2.5, 3.0}},
			[]string{"service", "status", "method", "uri"}),
	}

	reg.MustRegister(m.ReqByStatusCode, m.Duration)
	return m
}

func (m *CustomerMetrics) IncReqByStatusCode(status string) {
	m.ReqByStatusCode.With(prometheus.Labels{"service": m.service, "status": status}).Inc()
}

func (m *CustomerMetrics) MeasureDuration(start time.Time, method string, uri string, statusCode string) {
	m.Duration.With(prometheus.Labels{
		"service": m.service,
		"method":  method,
		"uri":     uri,
		"status":  statusCode,
	}).Observe(float64(time.Since(start).Seconds()))
}
