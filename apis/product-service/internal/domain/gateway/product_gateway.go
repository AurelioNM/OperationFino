package gateway

import (
	"cmd/product-service/internal/domain/entity"
	"context"
)

type ProductGateway interface {
	GetProductList(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, productID string) (*entity.Product, error)
	GetProductByName(ctx context.Context, productName string) (*entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (*string, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	DeleteProductByID(ctx context.Context, productID string) error
}
