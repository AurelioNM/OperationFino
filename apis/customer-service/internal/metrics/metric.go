package metrics

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

type CustomerMetrics struct {
	logger            slog.Logger
	Version           string
	AppInfo           *prometheus.GaugeVec
	TotalCustomers    prometheus.Gauge
	UpdateOnCustomers *prometheus.CounterVec
	Duration          *prometheus.HistogramVec
	SummaryDuration   prometheus.Summary
}

func NewCustomerMetrics(l slog.Logger, reg prometheus.Registerer) *CustomerMetrics {
	namespace := "OF_CustomerService"
	m := &CustomerMetrics{
		logger:  l,
		Version: "1.0",
		AppInfo: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "app_info",
				Help:      "Information about app environment"},
			[]string{"version"}),
		TotalCustomers: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "total_customers",
			Help:      "Number of customers on DB"}),
		UpdateOnCustomers: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "total_customers_update",
			Help:      "Number of upgraded devices"},
			[]string{"type"}),
		Duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "request_duration_seconds",
			Help:      "Duration of request",
			Buckets:   []float64{0.03, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.5, 0.6, 0.7, 0.8, 1.0, 1.5, 2.0, 2.5, 3.0}},
			[]string{"status", "method"}),
		SummaryDuration: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  namespace,
			Name:       "request_summary_duration_seconds",
			Help:       "Summary duration of request",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
	}

	m.AppInfo.With(prometheus.Labels{"version": m.Version}).Set(1)

	reg.MustRegister(m.TotalCustomers, m.AppInfo, m.UpdateOnCustomers, m.Duration)
	return m
}

func (m *CustomerMetrics) SetTotalCustomers(size int) {
	m.TotalCustomers.Set(float64(size))
}

func (m *CustomerMetrics) IncreaseTotalCustomers() {
	m.TotalCustomers.Inc()
}

func (m *CustomerMetrics) IncreaseUpdateOnCustomers() {
	m.UpdateOnCustomers.With(prometheus.Labels{"type": "router"}).Inc()
}
