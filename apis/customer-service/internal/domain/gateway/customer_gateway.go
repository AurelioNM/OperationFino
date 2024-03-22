package gateway

import (
	"cmd/customer-service/internal/domain/entity"
)

type CustomerGateway interface {
	GetCustomerList() ([]*entity.Customer, error)
	// GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error)
	// CreateCustomer(ctx context.Context, customer entity.Customer) error
	// UpdateCustomer(ctx context.Context, customer entity.Customer) error
	// DeleteCustomerByID(ctx context.Context, customerID string) error
}
