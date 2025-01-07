package gateway

import (
	"cmd/order-service/internal/domain/entity"
	"context"
)

type CustomerGateway interface {
	GetCustomerByEmail(ctx context.Context, customerEmail *string) (*entity.Customer, error)
}
