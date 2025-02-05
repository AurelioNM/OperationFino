package gateway

import (
	"cmd/customer-service/internal/domain/entity"
	"context"
)

type CustomerGateway interface {
	GetCustomerList(ctx context.Context) ([]*entity.Customer, error)
	GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error)
	GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error)
	GetCustomerByName(ctx context.Context, customerName string) (*entity.Customer, error)
	CreateCustomer(ctx context.Context, customer entity.Customer) (*string, error)
	UpdateCustomer(ctx context.Context, customer entity.Customer) error
	DeleteCustomerByID(ctx context.Context, customerID string) error
}
