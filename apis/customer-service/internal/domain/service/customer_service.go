package service

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"
	"cmd/customer-service/internal/resources/cache"
	"context"

	"log/slog"
)

type CustomerService interface {
	GetCustomerList(ctx context.Context) ([]*entity.Customer, error)
	GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error)
	V2GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error)
	GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error)
	V2GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error)
	CreateCustomer(ctx context.Context, customer entity.Customer) (*string, error)
	UpdateCustomer(ctx context.Context, customer entity.Customer) error
	DeleteCustomerByID(ctx context.Context, customerID string) error
}

type customerService struct {
	logger        slog.Logger
	customerGtw   gateway.CustomerGateway
	customerCache cache.CustomerCache
}

func NewCustomerService(l slog.Logger, g gateway.CustomerGateway, c cache.CustomerCache) CustomerService {
	return &customerService{
		logger:        *l.With("layer", "customer-service"),
		customerGtw:   g,
		customerCache: c,
	}
}

func (s *customerService) GetCustomerList(ctx context.Context) ([]*entity.Customer, error) {
	s.logger.Info("Getting all customers", "traceID", ctx.Value("traceID"))
	customerList, err := s.customerGtw.GetCustomerList(ctx)
	if err != nil {
		return nil, err
	}

	return customerList, nil
}

func (s *customerService) GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error) {
	s.logger.Info("Getting customer by ID", "ID", customerID, "traceID", ctx.Value("traceID"))
	customer, err := s.customerGtw.GetCustomerByID(ctx, customerID)
	if err != nil {
		s.logger.Error("Failed to get customer by ID", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return customer, nil
}

func (s *customerService) V2GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error) {
	s.logger.Info("Getting customer by ID", "ID", customerID, "traceID", ctx.Value("traceID"))
	customer, err := s.customerCache.ReadCacheByID(ctx, customerID)
	if err != nil {
		customer, err = s.customerGtw.GetCustomerByID(ctx, customerID)
		if err != nil {
			s.logger.Error("Failed to get customer by ID", "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		s.customerCache.WriteCacheByID(ctx, *customer)
		return customer, nil
	}

	return customer, nil
}

func (s *customerService) GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error) {
	s.logger.Info("Getting customer by email", "email", customerEmail, "traceID", ctx.Value("traceID"))
	customer, err := s.customerGtw.GetCustomerByEmail(ctx, customerEmail)
	if err != nil {
		s.logger.Error("Failed to get customer by email", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return customer, nil
}

func (s *customerService) V2GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error) {
	s.logger.Info("Getting customer by email", "email", customerEmail, "traceID", ctx.Value("traceID"))

	customer, err := s.customerCache.ReadCacheByEmail(ctx, customerEmail)
	if err != nil {
		customer, err = s.customerGtw.GetCustomerByEmail(ctx, customerEmail)
		if err != nil {
			s.logger.Error("Failed to get customer by email", "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		s.customerCache.WriteCacheByEmail(ctx, *customer)
		return customer, nil
	}

	return customer, nil
}

func (s *customerService) CreateCustomer(ctx context.Context, customer entity.Customer) (*string, error) {
	s.logger.Info("Creating new customer", "data", customer, "traceID", ctx.Value("traceID"))
	id, err := s.customerGtw.CreateCustomer(ctx, customer)
	if err != nil {
		s.logger.Error("Failed to create customer", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return id, nil
}

func (s *customerService) UpdateCustomer(ctx context.Context, customer entity.Customer) error {
	s.logger.Info("Updating customer", "data", customer)
	err := s.customerGtw.UpdateCustomer(ctx, customer)
	if err != nil {
		s.logger.Error("Failed to update customer by ID", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return nil
}

func (s *customerService) DeleteCustomerByID(ctx context.Context, customerID string) error {
	s.logger.Info("Deleting customer by ID", "ID", customerID)
	err := s.customerGtw.DeleteCustomerByID(ctx, customerID)
	if err != nil {
		s.logger.Error("Failed to delete customer by ID", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return nil
}
