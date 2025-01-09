package gateway

import (
	"cmd/order-service/internal/domain/entity"
	"context"
)

type OrderGateway interface {
	GetOrderByID(ctx context.Context, orderID *string) (*entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.Order) (*string, error)
}
