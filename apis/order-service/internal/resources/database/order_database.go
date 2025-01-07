package database

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"context"
	"database/sql"
	"log/slog"
)

type orderGateway struct {
	logger slog.Logger
	db     *sql.DB
}

func NewOrderGateway(l slog.Logger, db *sql.DB) gateway.OrderGateway {
	return &orderGateway{
		logger: *l.With("layer", "order-database"),
		db:     db,
	}
}

func (g *orderGateway) GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error) {
	g.logger.Debug("Getting order by ID from db", "ID", orderID, "traceID", ctx.Value("traceID"))
	return nil, nil
}

func (g *orderGateway) CreateOrder(ctx context.Context, order entity.Order) (*string, error) {
	g.logger.Debug("Inserting order into DB", "traceID", ctx.Value("traceID"))
	return nil, nil
}
