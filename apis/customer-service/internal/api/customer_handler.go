package api

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/service"
	"encoding/json"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/gorilla/mux"
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
	customerSvc service.CustomerService
}

func NewCustomerHandler(l slog.Logger, s service.CustomerService) CustomerHandler {
	return &customerHandler{
		logger:      *l.With("layer", "customer-handler"),
		customerSvc: s,
	}
}

func (h *customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Getting all customers")

	customers, err := h.customerSvc.GetCustomerList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(customers)
}

func (h *customerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("Getting customer by id", "id", id)
	customers, err := h.customerSvc.GetCustomerList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(customers)
	json.NewEncoder(w).Encode(customers[len(customers)-1])
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer entity.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Debug("Creating new customer")
	w.WriteHeader(http.StatusCreated)
}

func (h *customerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var customer entity.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("Updating customer", "id", id, "body", customer)
	w.WriteHeader(http.StatusOK)
}

func (h *customerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("Deleting customer", "id", id)
	w.WriteHeader(http.StatusNoContent)
}
