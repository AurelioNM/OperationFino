package api

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/metrics"
	"encoding/json"
	"fmt"
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

type response struct {
	Message     string                 `json:"message"`
	Timestamp   time.Time              `json:"timestamp"`
	ElapsedTime string                 `json:"elapsed_time"`
	Data        map[string]interface{} `json:"data"`
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
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "getAll", time.Since(now))
		return
	}

	h.metrics.SetTotalCustomers(len(customers))
	h.metrics.IncreaseRequestsByStatusCode("200")
	h.metrics.Duration.With(prometheus.Labels{"method": "GET", "uri": "/v1/customers", "status": "200"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.SummaryDuration.Observe(time.Since(now).Seconds())

	h.buildResponse(w, "All customers", time.Since(now), map[string]interface{}{"page": customers, "page_size": len(customers)})
}

func (h *customerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("GET customer by ID request")
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := h.customerSvc.GetCustomerByID(id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "getByID", time.Since(now))
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "GET", "uri": "/v1/customers/{customerId}", "status": "200"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseRequestsByStatusCode("200")

	h.buildResponse(w, fmt.Sprintf("Customer by ID: %s", id), time.Since(now), map[string]interface{}{"customer": customer})
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("POST customer request")
	var customer entity.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "create", time.Since(now))
		return
	}

	id, err := h.customerSvc.CreateCustomer(customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "create", time.Since(now))
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "POST", "uri": "/v1/customers", "status": "201"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseTotalCustomers()
	h.metrics.IncreaseRequestsByStatusCode("201")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	h.buildResponse(w, "Customer created", time.Since(now), map[string]interface{}{"id": id})
}

func (h *customerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("PUT customer by ID request")
	vars := mux.Vars(r)
	id := vars["id"]
	var customer entity.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "update", time.Since(now))
		return
	}
	customer.ID = &id

	err = h.customerSvc.UpdateCustomer(customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "update", time.Since(now))
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "PUT", "uri": "/v1/customers/{customerId}", "status": "200"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseRequestsByStatusCode("200")

	h.buildResponse(w, "Customer updated", time.Since(now), map[string]interface{}{"id": id})
}

func (h *customerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	h.logger.Debug("DELETE customer by ID request")
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.customerSvc.DeleteCustomerByID(id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "delete", time.Since(now))
		return
	}

	h.metrics.Duration.With(prometheus.Labels{"method": "DELETE", "uri": "/v1/customers/{customerId}", "status": "204"}).Observe(float64(time.Since(now).Seconds()))
	h.metrics.IncreaseRequestsByStatusCode("204")

	h.buildResponse(w, "Customer deleted", time.Since(now), map[string]interface{}{})
}

func (h *customerHandler) buildResponse(w http.ResponseWriter, message string, elapsedTime time.Duration, data map[string]interface{}) {
	res := &response{
		Message:     message,
		Timestamp:   time.Now(),
		ElapsedTime: fmt.Sprintf("%dms", elapsedTime.Milliseconds()),
		Data:        data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *customerHandler) buildErrorResponse(w http.ResponseWriter, error string, statusCode int, operation string, elapsed time.Duration) {
	h.metrics.IncreaseRequestsByStatusCode(fmt.Sprint(statusCode))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	h.buildResponse(w, fmt.Sprintf("Error on %s customer: %s", operation, error), elapsed, map[string]interface{}{})
}
