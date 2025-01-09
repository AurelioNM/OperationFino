package gateway

import (
	"cmd/order-service/internal/resources/client/dto"
	"context"
)

type ProductGateway interface {
	GetProductByName(ctx context.Context, productName *string) (*dto.Product, error)
}
