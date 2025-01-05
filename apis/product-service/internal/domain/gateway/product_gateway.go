package gateway

import (
	"cmd/product-service/internal/domain/entity"
	"context"
)

type ProductGateway interface {
	GetProductList(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, productID string) (*entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (*string, error)
}
