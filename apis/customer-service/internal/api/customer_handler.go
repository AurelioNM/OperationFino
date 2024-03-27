package api

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/service"
	"encoding/json"
	"net/http"

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
	customers, err := h.customerSvc.GetCustomerList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(customers)
}

func (h *customerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := h.customerSvc.GetCustomerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}

func (h *customerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *customerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.customerSvc.DeleteCustomerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
