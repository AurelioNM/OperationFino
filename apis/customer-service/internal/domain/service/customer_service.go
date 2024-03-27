package service

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"

	"log/slog"
)

type CustomerService interface {
	GetCustomerList() ([]*entity.Customer, error)
	GetCustomerByID(customerID string) (*entity.Customer, error)
	CreateCustomer(customer entity.Customer) (*string, error)
	UpdateCustomer(customer entity.Customer) error
	DeleteCustomerByID(customerID string) error
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

func (s *customerService) GetCustomerByID(customerID string) (*entity.Customer, error) {
	s.logger.Debug("Getting customer by ID", "ID", customerID)
	customer, err := s.customerGtw.GetCustomerByID(customerID)
	if err != nil {
		s.logger.Error("Failed to get customer by ID", "error", err)
		return nil, err
	}

	return customer, nil
}

func (s *customerService) CreateCustomer(customer entity.Customer) (*string, error) {
	s.logger.Info("Creating new customer", "data", customer)
	id, err := s.customerGtw.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *customerService) UpdateCustomer(customer entity.Customer) error {
	s.logger.Info("Updating customer", "data", customer)
	err := s.customerGtw.UpdateCustomer(customer)
	if err != nil {
		s.logger.Error("Failed to update customer by ID", "error", err)
		return err
	}

	return nil
}

func (s *customerService) DeleteCustomerByID(customerID string) error {
	s.logger.Info("Deleting customer by ID", "ID", customerID)
	err := s.customerGtw.DeleteCustomerByID(customerID)
	if err != nil {
		s.logger.Error("Failed to delete customer by ID", "error", err)
		return err
	}

	return nil
}
