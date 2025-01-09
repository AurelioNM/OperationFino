package service

import (
	"cmd/product-service/internal/domain/entity"
	"cmd/product-service/internal/domain/gateway"
	"context"
	"log/slog"
)

type ProductService interface {
	GetProductList(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, productID string) (*entity.Product, error)
	GetProductByName(ctx context.Context, productName string) (*entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (*string, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	DeleteProductByID(ctx context.Context, productID string) error
}

type productService struct {
	logger     slog.Logger
	productGtw gateway.ProductGateway
}

func NewProductService(l slog.Logger, g gateway.ProductGateway) ProductService {
	return &productService{
		logger:     *l.With("layer", "product-service"),
		productGtw: g,
	}
}

func (s *productService) GetProductList(ctx context.Context) ([]*entity.Product, error) {
	s.logger.Info("Getting all products", "traceID", ctx.Value("traceID"))
	productList, err := s.productGtw.GetProductList(ctx)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (s *productService) GetProductByID(ctx context.Context, productID string) (*entity.Product, error) {
	s.logger.Info("Getting product by ID", "ID", productID, "traceID", ctx.Value("traceID"))
	product, err := s.productGtw.GetProductByID(ctx, productID)
	if err != nil {
		s.logger.Error("Failed to get product by ID", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProductByName(ctx context.Context, productName string) (*entity.Product, error) {
	s.logger.Info("Getting product by name", "productName", productName, "traceID", ctx.Value("traceID"))
	product, err := s.productGtw.GetProductByName(ctx, productName)
	if err != nil {
		s.logger.Error("Failed to get product by name", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return product, nil
}

func (s *productService) CreateProduct(ctx context.Context, product entity.Product) (*string, error) {
	s.logger.Info("Creating new product", "data", product, "traceID", ctx.Value("traceID"))
	id, err := s.productGtw.CreateProduct(ctx, product)
	if err != nil {
		s.logger.Error("Failed to create product", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return id, nil
}

func (s *productService) UpdateProduct(ctx context.Context, product entity.Product) error {
	s.logger.Info("Updating product", "data", product)
	err := s.productGtw.UpdateProduct(ctx, product)
	if err != nil {
		s.logger.Error("Failed to update product by ID", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return nil
}

func (s *productService) DeleteProductByID(ctx context.Context, productID string) error {
	s.logger.Info("Deleting product by ID", "ID", productID)
	err := s.productGtw.DeleteProductByID(ctx, productID)
	if err != nil {
		s.logger.Error("Failed to delete product by ID", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return nil
}
