package api

import (
	"cmd/customer-service/internal/domain/entity"
	"encoding/json"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/gorilla/mux"
)

type CustomerHandler interface {
	GetCustomers(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	logger    *slog.Logger
	Customers []entity.Customer
}

func NewCustomerHandler(logger *slog.Logger) CustomerHandler {
	return &customerHandler{
		logger: logger,
		Customers: []entity.Customer{
			{ID: 1, Name: "Fulano", Surname: "Beltrano", Email: "fulano@gmail.com"},
			{ID: 2, Name: "Ciclano", Surname: "Nunes", Email: "ciclano@gmail.com"},
		},
	}
}

func (h *customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Retrieving all customers")
	json.NewEncoder(w).Encode(h.Customers)
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer entity.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Customers = append(h.Customers, customer)
	h.logger.Debug("Creating new customer", "len", len(h.Customers))

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
