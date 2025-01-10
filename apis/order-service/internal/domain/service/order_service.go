package service

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"cmd/order-service/internal/resources/client/dto"
	"context"
	"log/slog"
	"time"

	"github.com/oklog/ulid/v2"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error)
	GetOrdersByCustomerID(ctx context.Context, customerID string) ([]*entity.Order, error)
	CreateOrder(ctx context.Context, orderRequest *entity.OrderRequest) (*string, error)
	DeleteOrderByID(ctx context.Context, orderID string) error
}

type orderService struct {
	logger      slog.Logger
	orderGtw    gateway.OrderGateway
	customerGtw gateway.CustomerGateway
	productGtw  gateway.ProductGateway
}

func NewOrderService(l slog.Logger, g gateway.OrderGateway, c gateway.CustomerGateway, p gateway.ProductGateway) OrderService {
	return &orderService{
		logger:      *l.With("layer", "order-service"),
		orderGtw:    g,
		customerGtw: c,
		productGtw:  p,
	}
}

func (s *orderService) GetOrderByID(ctx context.Context, orderID string) (*entity.Order, error) {
	s.logger.Info("Getting order by ID", "ID", orderID, "traceID", ctx.Value("traceID"))
	order, err := s.orderGtw.GetOrderByID(ctx, &orderID)
	if err != nil {
		s.logger.Error("Failed to get order by ID", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetOrdersByCustomerID(ctx context.Context, customerID string) ([]*entity.Order, error) {
	s.logger.Info("Getting orders list by customerID", "customerID", customerID, "traceID", ctx.Value("traceID"))
	orders, err := s.orderGtw.GetOrdersByCustomerID(ctx, &customerID)
	if err != nil {
		s.logger.Error("Failed to get orders list by customerID", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return orders, nil
}

func (s *orderService) CreateOrder(ctx context.Context, orderRequest *entity.OrderRequest) (*string, error) {
	s.logger.Info("Getting customer and product info to build Order", "orderRequest", orderRequest, "traceID", ctx.Value("traceID"))

	customer, err := s.customerGtw.GetCustomerByEmail(ctx, &orderRequest.CustomerEmail)
	if err != nil {
		s.logger.Error("Failed to get customer", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	products := make([]dto.Product, 0)
	for _, productRequest := range orderRequest.Products {
		product, err := s.productGtw.GetProductByName(ctx, &productRequest.Name)
		if err != nil {
			s.logger.Error("Failed to get product", "productName", productRequest.Name, "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		s.logger.Debug("Inserting product on list", "traceID", ctx.Value("traceID"))
		products = append(products, *product)
	}

	s.logger.Debug("Building order", "traceID", ctx.Value("traceID"))
	orderID := ulid.Make().String()
	order := &entity.Order{
		ID:        &orderID,
		Customer:  *customer,
		Products:  products,
		CreatedAt: time.Now(),
	}

	s.logger.Info("Creating new order", "order", *order, "traceID", ctx.Value("traceID"))
	id, err := s.orderGtw.CreateOrder(ctx, order)
	if err != nil {
		s.logger.Error("Failed to create order", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return id, nil
}

func (s *orderService) DeleteOrderByID(ctx context.Context, orderID string) error {
	s.logger.Info("Deleting order by ID", "ID", orderID, "traceID", ctx.Value("traceID"))
	err := s.orderGtw.DeleteOrderByID(ctx, &orderID)
	if err != nil {
		s.logger.Error("Failed to delete order by ID", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return nil
}
