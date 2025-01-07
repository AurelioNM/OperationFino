package gateway

import (
	"cmd/order-service/internal/domain/entity"
	"context"
)

type ProductGateway interface {
	GetProductByName(ctx context.Context, productName *string) (*entity.Product, error)
}
