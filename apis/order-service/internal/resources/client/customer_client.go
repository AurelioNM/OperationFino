package client

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"cmd/order-service/internal/metrics"
	"context"
	"log/slog"
)

type customerGateway struct {
	logger  slog.Logger
	metrics metrics.OrderMetrics
}

func NewCustomerGateway(l slog.Logger, m metrics.OrderMetrics) gateway.CustomerGateway {
	return &customerGateway{
		logger:  *l.With("layer", "customer-client"),
		metrics: m,
	}
}

func (g *customerGateway) GetCustomerByEmail(ctx context.Context, customerEmail *string) (*entity.Customer, error) {
	g.logger.Debug("Calling customer-service to get on getCustomerByEmail", "customerEmail", customerEmail, "traceID", ctx.Value("traceID"))
	return nil, nil
}
