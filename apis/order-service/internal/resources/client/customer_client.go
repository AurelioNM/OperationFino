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

type customerGateway struct {
	logger  slog.Logger
	metrics *metrics.OrderMetrics
}

func NewCustomerGateway(l slog.Logger, m *metrics.OrderMetrics) gateway.CustomerGateway {
	return &customerGateway{
		logger:  *l.With("layer", "customer-client"),
		metrics: m,
	}
}

func (g *customerGateway) GetCustomerByEmail(ctx context.Context, customerEmail *string) (*dto.Customer, error) {
	g.logger.Info("Calling customer-service to get getCustomerByEmail", "customerEmail", customerEmail, "traceID", ctx.Value("traceID"))
	url := fmt.Sprintf("http://of-customer-service:8001/v2/customers/email/%s", *customerEmail)
	now := time.Now()

	res, err := http.Get(url)
	if err != nil {
		g.logger.Error("Customer-service request failed", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}
	g.metrics.MeasureExternalDuration(now, "customer-service", "GET", "/v2/customers/email/{email}", "200")

	body, err := g.getBodyFromResponse(ctx, res)
	if err != nil {
		return nil, err
	}
	g.logger.Debug("Customer-service body data", "body", string(body), "traceID", ctx.Value("traceID"))

	var responseDTO dto.GetCustomerByEmailResponseDTO
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		g.logger.Error("Failed to unmarshal customer response body", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	g.logger.Info("Customer-service request successfull", "responseDTO", responseDTO, "traceID", ctx.Value("traceID"))
	return &responseDTO.Data.Customer, nil
}

func (g *customerGateway) getBodyFromResponse(ctx context.Context, res *http.Response) ([]byte, error) {
	if res.StatusCode != http.StatusOK {
		g.logger.Error("Request status code is not OK", "statusCode", res.StatusCode, "traceID", ctx.Value("traceID"))
		return nil, fmt.Errorf("Got statusCode %d from customer request", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		g.logger.Error("Failed to read customer response body", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return body, nil
}
