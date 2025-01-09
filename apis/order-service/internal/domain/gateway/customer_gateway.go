package gateway

import (
	"cmd/order-service/internal/resources/client/dto"
	"context"
)

type CustomerGateway interface {
	GetCustomerByEmail(ctx context.Context, customerEmail *string) (*dto.Customer, error)
}
