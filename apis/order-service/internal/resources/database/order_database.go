package database

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"context"
	"log/slog"
)

type orderGateway struct {
	logger slog.Logger
}

func NewOrderGateway(l slog.Logger) gateway.OrderGateway {
	return &orderGateway{
		logger: *l.With("layer", "order-database"),
	}
}

func (g *orderGateway) GetOrderByID(ctx context.Context, orderID *string) (*entity.Order, error) {
	g.logger.Debug("Getting order by ID from db", "ID", orderID, "traceID", ctx.Value("traceID"))
	return nil, nil
}

func (g *orderGateway) CreateOrder(ctx context.Context, order *entity.Order) (*string, error) {
	g.logger.Debug("Inserting order into DB", "traceID", ctx.Value("traceID"))
	return order.ID, nil
}
