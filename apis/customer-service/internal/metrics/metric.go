package metrics

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

type CustomerMetrics struct {
	logger         slog.Logger
	Version        string
	AppInfo        *prometheus.GaugeVec
	TotalCustomers prometheus.Gauge
}

type Device struct {
	ID       string `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
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
	}

	m.AppInfo.With(prometheus.Labels{"version": m.Version}).Set(1)

	reg.MustRegister(m.TotalCustomers, m.AppInfo)
	return m
}

func (m *CustomerMetrics) SetTotalCustomers(size int) {
	m.TotalCustomers.Set(float64(size))
}

func (m *CustomerMetrics) IncreaseTotalCustomers() {
	m.TotalCustomers.Inc()
}
