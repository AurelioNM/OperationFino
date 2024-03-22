package service

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"

	"log/slog"
)

type CustomerService interface {
	GetCustomerList() ([]*entity.Customer, error)
	// GetCustomerByID(customerID string) (*entity.Customer, error)
	// CreateCustomer(customer entity.Customer) error
	// UpdateCustomer(customer entity.Customer) error
	// DeleteCustomerByID(customerID string) error
}

type customerService struct {
	logger      slog.Logger
	customerGtw gateway.CustomerGateway
}

func NewCustomerService(l slog.Logger, c gateway.CustomerGateway) CustomerService {
	return &customerService{
		logger:      *l.With("layer", "customer-service"),
		customerGtw: c,
	}
}

func (s *customerService) GetCustomerList() ([]*entity.Customer, error) {
	s.logger.Info("Getting all customers")
	customerList, err := s.customerGtw.GetCustomerList()
	if err != nil {
		return nil, err
	}

	return customerList, nil
}
