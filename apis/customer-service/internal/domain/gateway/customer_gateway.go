package gateway

import (
	"cmd/customer-service/internal/domain/entity"
)

type CustomerGateway interface {
	GetCustomerList() ([]*entity.Customer, error)
	GetCustomerByID(customerID string) (*entity.Customer, error)
	CreateCustomer(customer entity.Customer) (*string, error)
	UpdateCustomer(customer entity.Customer) error
	DeleteCustomerByID(customerID string) error
}
