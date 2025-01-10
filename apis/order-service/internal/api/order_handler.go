package api

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/service"
	"cmd/order-service/internal/metrics"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

type OrderHandler interface {
	GetOrderByID(w http.ResponseWriter, r *http.Request)
	GetOrdersByCustomerID(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrderByID(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	logger   slog.Logger
	metrics  *metrics.OrderMetrics
	orderSvc service.OrderService
}

type response struct {
	Message     string                 `json:"message"`
	Timestamp   time.Time              `json:"timestamp"`
	ElapsedTime string                 `json:"elapsed_time"`
	Data        map[string]interface{} `json:"data"`
}

func NewOrderHandler(l slog.Logger, m *metrics.OrderMetrics, s service.OrderService) OrderHandler {
	return &orderHandler{
		logger:   *l.With("layer", "order-handler"),
		metrics:  m,
		orderSvc: s,
	}
}

func (h *orderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("GET order by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	order, err := h.orderSvc.GetOrderByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "GET", "/v1/orders/{orderId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v1/orders/{orderId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, fmt.Sprintf("Order by ID: %s", id), now, map[string]interface{}{"order": order})
}

func (h *orderHandler) GetOrdersByCustomerID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("GET orders by customerID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	customerID := vars["customerID"]
	h.logger.Debug("Vars", "customerID", customerID, "traceID", ctx.Value("traceID"))

	orders, err := h.orderSvc.GetOrdersByCustomerID(ctx, customerID)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "GET", "/v1/orders/customers/{customerId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v1/orders/customers/{customerId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, fmt.Sprintf("Order by ID: %s", customerID), now, map[string]interface{}{"orders": orders})
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("POST order request", "traceID", ctx.Value("traceID"))

	var order *entity.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "POST", "/v1/orders", now)
		return
	}

	id, err := h.orderSvc.CreateOrder(ctx, order)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "POST", "/v1/orders", now)
		return
	}

	h.metrics.MeasureDuration(now, "POST", "/v1/orders", "201")
	h.metrics.IncReqByStatusCode("201")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	h.buildResponse(w, "Order created", now, map[string]interface{}{"id": id})
}

func (h *orderHandler) DeleteOrderByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("DELETE order by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.orderSvc.DeleteOrderByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "DELETE", "/v1/orders/{orderId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "DELETE", "/v1/orders/{orderId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "Product deleted", now, map[string]interface{}{})
}

func (h *orderHandler) getContext(r *http.Request) context.Context {
	traceID := r.Header.Get("X-Trace-ID")
	if traceID == "" {
		traceID = ulid.Make().String()
	}

	ctx := context.WithValue(r.Context(), "traceID", traceID)
	return ctx
}

func (h *orderHandler) buildErrorResponse(w http.ResponseWriter, error string, statusCode int, method string, uri string, start time.Time) {
	h.metrics.MeasureDuration(start, method, uri, fmt.Sprint(statusCode))
	h.metrics.IncReqByStatusCode(fmt.Sprint(statusCode))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	h.buildResponse(w, fmt.Sprintf("Error on %s order: %s", method, error), start, map[string]interface{}{})
}

func (h *orderHandler) buildResponse(w http.ResponseWriter, message string, start time.Time, data map[string]interface{}) {
	res := &response{
		Message:     message,
		Timestamp:   time.Now(),
		ElapsedTime: fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		Data:        data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
