package service

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"context"
	"log/slog"
	"time"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error)
	CreateOrder(ctx context.Context, orderRequest *entity.OrderRequest) (*string, error)
}

type orderService struct {
	logger      slog.Logger
	orderGtw    gateway.OrderGateway
	customerGtw gateway.CustomerGateway
	productGtw  gateway.ProductGateway
}

func NewOrderService(l slog.Logger, g gateway.OrderGateway) OrderService {
	return &orderService{
		logger:   *l.With("layer", "order-service"),
		orderGtw: g,
	}
}

func (s *orderService) GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error) {
	s.logger.Info("Getting order by ID", "ID", orderID, "traceID", ctx.Value("traceID"))
	order, err := s.orderGtw.GetOrderByID(ctx, orderID)
	if err != nil {
		s.logger.Error("Failed to get order by ID", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return order, nil
}

func (s *orderService) CreateOrder(ctx context.Context, orderRequest *entity.OrderRequest) (*string, error) {
	s.logger.Info("Getting customer and product info to build Order", "orderRequest", orderRequest, "traceID", ctx.Value("traceID"))

	customer, err := s.customerGtw.GetCustomerByEmail(ctx, &orderRequest.CustomerEmail)
	if err != nil {
		s.logger.Error("Failed to get customer", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	products := make([]entity.Product, 0)
	for _, productRequest := range orderRequest.Products {
		product, err := s.productGtw.GetProductByName(ctx, &productRequest.Name)
		if err != nil {
			s.logger.Error("Failed to get product", "productName", productRequest.Name, "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		products = append(products, *product)
	}

	order := &entity.Order{
		Customer:  *customer,
		Products:  products,
		CreatedAt: time.Now(),
	}

	s.logger.Info("Creating new order", "data", orderRequest, "traceID", ctx.Value("traceID"))
	id, err := s.orderGtw.CreateOrder(ctx, order)
	if err != nil {
		s.logger.Error("Failed to create order", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return id, nil
}
