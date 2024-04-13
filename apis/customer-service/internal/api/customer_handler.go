package api

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/metrics"
	"encoding/json"
	"net/http"
	"time"

	"log/slog"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type CustomerHandler interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	GetCustomerByID(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	logger      slog.Logger
	metrics     *metrics.CustomerMetrics
	customerSvc service.CustomerService
}

func NewCustomerHandler(l slog.Logger, m *metrics.CustomerMetrics, s service.CustomerService) CustomerHandler {
	return &customerHandler{
		logger:      *l.With("layer", "customer-handler"),
		metrics:     m,
		customerSvc: s,
	}
}

func (h *customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("GET customers request")
	customers, err := h.customerSvc.GetCustomerList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.metrics.SetTotalCustomers(len(customers))
	h.metrics.IncreaseRequestsByStatusCode("200")
	h.metrics.Duration.With(prometheus.Labels{"method": "GET", "uri": "/v1/customers", "status": "200"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.SummaryDuration.Observe(time.Since(now).Seconds())

	json.NewEncoder(w).Encode(customers)
}

func (h *customerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("GET customer by ID request")
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := h.customerSvc.GetCustomerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "GET", "uri": "/v1/customers/{customerId}", "status": "200"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseRequestsByStatusCode("200")
	json.NewEncoder(w).Encode(customer)
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("POST customer request")
	var customer entity.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.logger.Error("Failed to create customer", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.customerSvc.CreateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseJson, err := json.Marshal(map[string]*string{"id": id})
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusBadGateway)
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "POST", "uri": "/v1/customers", "status": "201"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseTotalCustomers()
	h.metrics.IncreaseRequestsByStatusCode("201")

	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}

func (h *customerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("PUT customer by ID request")
	vars := mux.Vars(r)
	id := vars["id"]
	var customer entity.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer.ID = &id

	err = h.customerSvc.UpdateCustomer(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		h.metrics.IncreaseRequestsByStatusCode("404")
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "PUT", "uri": "/v1/customers/{customerId}", "status": "200"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseRequestsByStatusCode("200")

	w.WriteHeader(http.StatusOK)
}

func (h *customerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("DELETE customer by ID request")
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.customerSvc.DeleteCustomerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "DELETE", "uri": "/v1/customers/{customerId}", "status": "204"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseRequestsByStatusCode("204")
	w.WriteHeader(http.StatusNoContent)
}
