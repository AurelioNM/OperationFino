package client

import (
	"cmd/order-service/internal/domain/gateway"
	"cmd/order-service/internal/metrics"
	"cmd/order-service/internal/resources/client/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type productGateway struct {
	logger  slog.Logger
	metrics *metrics.OrderMetrics
}

func NewProductGateway(l slog.Logger, m *metrics.OrderMetrics) gateway.ProductGateway {
	return &productGateway{
		logger:  *l.With("layer", "product-client"),
		metrics: m,
	}
}

func (g *productGateway) GetProductByName(ctx context.Context, productName *string) (*dto.Product, error) {
	g.logger.Info("Calling product-service to get on getProductByName", "productName", productName, "traceID", ctx.Value("traceID"))
	url := fmt.Sprintf("http://of-product-service:8002/v1/products/name/%s", *productName)
	start := time.Now()

	res, err := http.Get(url)
	g.metrics.MeasureExternalDuration(start, "product-service", "GET", "/v1/products/name/{name}", "")
	if err != nil {
		g.logger.Error("Product-service request failed", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	body, err := g.getBodyFromResponse(ctx, res)
	if err != nil {
		return nil, err
	}
	g.logger.Debug("Product-service body data", "body", string(body), "traceID", ctx.Value("traceID"))

	var responseDTO dto.GetProductByNameResponseDTO
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		g.logger.Error("Failed to unmarshal product response body", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	g.logger.Info("Product-service request successfull", "responseDTO", responseDTO, "traceID", ctx.Value("traceID"))
	return &responseDTO.Data.Product, nil
}

func (g *productGateway) getBodyFromResponse(ctx context.Context, res *http.Response) ([]byte, error) {
	if res.StatusCode != http.StatusOK {
		g.logger.Error("Request status code is not OK", "statusCode", res.StatusCode, "traceID", ctx.Value("traceID"))
		return nil, fmt.Errorf("Got statusCode %d from product request", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		g.logger.Error("Failed to read product response body", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return body, nil
}
