package api

import (
	"cmd/product-service/internal/domain/entity"
	"cmd/product-service/internal/domain/service"
	"cmd/product-service/internal/metrics"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

type ProductHandler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductByID(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	logger     slog.Logger
	metrics    *metrics.ProductMetrics
	productSvc service.ProductService
}

type response struct {
	Message     string                 `json:"message"`
	Timestamp   time.Time              `json:"timestamp"`
	ElapsedTime string                 `json:"elapsed_time"`
	Data        map[string]interface{} `json:"data"`
}

func NewProductHandler(l slog.Logger, m *metrics.ProductMetrics, s service.ProductService) ProductHandler {
	return &productHandler{
		logger:     *l.With("layer", "product-handler"),
		metrics:    m,
		productSvc: s,
	}
}

func (h *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("GET all products request", "traceID", ctx.Value("traceID"))

	products, err := h.productSvc.GetProductList(ctx)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "GET", "/v1/products", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v1/products", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "All products", now, map[string]interface{}{"page_size": len(products), "page_content": products})
}

func (h *productHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("GET product by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	product, err := h.productSvc.GetProductByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "GET", "/v1/products/{productId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "GET", "/v1/products/{productId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, fmt.Sprintf("Product by ID: %s", id), now, map[string]interface{}{"product": product})
}

func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("POST product request", "traceID", ctx.Value("traceID"))

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "POST", "/v1/products", now)
		return
	}

	id, err := h.productSvc.CreateProduct(ctx, product)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "POST", "/v1/products", now)
		return
	}

	h.metrics.MeasureDuration(now, "POST", "/v1/products", "201")
	h.metrics.IncReqByStatusCode("201")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	h.buildResponse(w, "Product created", now, map[string]interface{}{"id": id})
}

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("PUT product by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]
	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusBadRequest, "PUT", "/v1/products/{productId}", now)
		return
	}
	product.ID = &id

	err = h.productSvc.UpdateProduct(ctx, product)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "PUT", "/v1/products/{productId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "PUT", "/v1/products/{productId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "Product updated", now, map[string]interface{}{"id": id})
}

func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := h.getContext(r)
	h.logger.Debug("DELETE product by ID request", "traceID", ctx.Value("traceID"))

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.productSvc.DeleteProductByID(ctx, id)
	if err != nil {
		h.buildErrorResponse(w, err.Error(), http.StatusNotFound, "DELETE", "/v1/products/{productId}", now)
		return
	}

	h.metrics.MeasureDuration(now, "DELETE", "/v1/products/{productId}", "200")
	h.metrics.IncReqByStatusCode("200")

	h.buildResponse(w, "Product deleted", now, map[string]interface{}{})
}

func (h *productHandler) getContext(r *http.Request) context.Context {
	traceID := r.Header.Get("X-Trace-ID")
	if traceID == "" {
		traceID = ulid.Make().String()
	}

	ctx := context.WithValue(r.Context(), "traceID", traceID)
	return ctx
}

func (h *productHandler) buildErrorResponse(w http.ResponseWriter, error string, statusCode int, method string, uri string, start time.Time) {
	h.metrics.MeasureDuration(start, method, uri, fmt.Sprint(statusCode))
	h.metrics.IncReqByStatusCode(fmt.Sprint(statusCode))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	h.buildResponse(w, fmt.Sprintf("Error on %s product: %s", method, error), start, map[string]interface{}{})
}

func (h *productHandler) buildResponse(w http.ResponseWriter, message string, start time.Time, data map[string]interface{}) {
	res := &response{
		Message:     message,
		Timestamp:   time.Now(),
		ElapsedTime: fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		Data:        data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
