package client

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"cmd/order-service/internal/metrics"
	"context"
	"log/slog"
)

type productGateway struct {
	logger  slog.Logger
	metrics metrics.OrderMetrics
}

func NewProductGateway(l slog.Logger, m metrics.OrderMetrics) gateway.ProductGateway {
	return &productGateway{
		logger:  *l.With("layer", "product-client"),
		metrics: m,
	}
}

func (g *productGateway) GetProductByName(ctx context.Context, productName *string) (*entity.Product, error) {
	g.logger.Debug("Calling product-service to get on getProductByName", "productName", productName, "traceID", ctx.Value("traceID"))
	return nil, nil
}
