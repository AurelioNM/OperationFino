package api

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/metrics"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"log/slog"

	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

type CustomerHandler interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	GetCustomerByID(w http.ResponseWriter, r *http.Request)
	V2GetCustomerByID(w http.ResponseWriter, r *http.Request)
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
	ctx := h.getContext(r)
	h.logger.Debug("GET all customers request", "traceID", ctx.Value("traceID"))

	customers, err := h.customerSvc.GetCustomerList(ctx)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "GET", "/v1/customers", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v1/customers", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "All customers", now, map[string]interface{}{"page": customers, "page_size": len(customers)})
}

func (h *customerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("GET customer by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := h.customerSvc.V2GetCustomerByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "GET", "/v1/customers/{customerId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v1/customers/{customerId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, fmt.Sprintf("Customer by ID: %s", id), now, map[string]interface{}{"customer": customer})
}

func (h *customerHandler) V2GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("GET customer by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := h.customerSvc.V2GetCustomerByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "GET", "/v2/customers/{customerId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v2/customers/{customerId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, fmt.Sprintf("Customer by ID: %s", id), now, map[string]interface{}{"customer": customer})
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("POST customer request", "traceID", ctx.Value("traceID"))

	var customer entity.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "POST", "/v1/customers", now)
		return
	}

	id, err := h.customerSvc.CreateCustomer(ctx, customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "POST", "/v1/customers", now)
		return
	}

	h.metrics.MeasureDuration(now, "POST", "/v1/customers", "201")
	h.metrics.IncReqByStatusCode("201")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	h.buildResponse(w, "Customer created", now, map[string]interface{}{"id": id})
}

func (h *customerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("PUT customer by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]
	var customer entity.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "PUT", "/v1/customers", now)
		return
	}
	customer.ID = &id

	err = h.customerSvc.UpdateCustomer(ctx, customer)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "PUT", "/v1/customers", now)
		return
	}

	h.metrics.MeasureDuration(now, "PUT", "/v1/customers", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "Customer updated", now, map[string]interface{}{"id": id})
}

func (h *customerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("DELETE customer by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.customerSvc.DeleteCustomerByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "DELETE", "/v1/customers/{customerId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "DELETE", "/v1/customers/{customerId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "Customer deleted", now, map[string]interface{}{})
}

func (h *customerHandler) getContext(r *http.Request) context.Context {
	traceID := r.Header.Get("X-Trace-ID")
	if traceID == "" {
		traceID = ulid.Make().String()
	}

	ctx := context.WithValue(r.Context(), "traceID", traceID)
	return ctx
}

func (h *customerHandler) buildResponse(w http.ResponseWriter, message string, start time.Time, data map[string]interface{}) {
	res := &response{
		Message:     message,
		Timestamp:   time.Now(),
		ElapsedTime: fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		Data:        data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *customerHandler) buildErrorResponse(w http.ResponseWriter, error string, statusCode int, method string, uri string, start time.Time) {
	h.metrics.MeasureDuration(start, method, uri, fmt.Sprint(statusCode))
	h.metrics.IncReqByStatusCode(fmt.Sprint(statusCode))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	h.buildResponse(w, fmt.Sprintf("Error on %s customer: %s", method, error), start, map[string]interface{}{})
}
