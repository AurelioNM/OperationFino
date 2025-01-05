package database

import (
	"cmd/product-service/internal/domain/entity"
	"cmd/product-service/internal/domain/gateway"
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/oklog/ulid/v2"
)

type productGateway struct {
	logger slog.Logger
	db     *sql.DB
}

func NewProductGateway(l slog.Logger, db *sql.DB) gateway.ProductGateway {
	return &productGateway{
		logger: *l.With("layer", "product-database"),
		db:     db,
	}
}

func (g *productGateway) GetProductList(ctx context.Context) ([]*entity.Product, error) {
	g.logger.Debug("Getting all products from DB", "traceID", ctx.Value("traceID"))
	query := "SELECT product_id, name, description, price, quantity, created_at, updated_at FROM products;"

	rows, err := g.db.Query(query)
	if err != nil {
		g.logger.Error("Failed to get products from db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer rows.Close()
	products := make([]*entity.Product, 0)
	for rows.Next() {
		product := &entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			g.logger.Error("Error scaning row", "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		products = append(products, product)
	}

	g.logger.Info("Found product list on DB", "size", len(products))
	return products, nil
}

func (g *productGateway) GetProductByID(ctx context.Context, productID string) (*entity.Product, error) {
	g.logger.Debug("Getting product by ID from db", "ID", productID, "traceID", ctx.Value("traceID"))
	query := "SELECT product_id, name, description, price, quantity, created_at, updated_at FROM products WHERE product_id = $1;"

	rows, err := g.db.Query(query, productID)
	if err != nil {
		g.logger.Error("Failed to get product by ID from db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			g.logger.Error("Error scaning product row", "error", err)
			return nil, err
		}
		return &product, nil
	}

	return nil, fmt.Errorf("No product found with ID=%s", productID)
}

func (g *productGateway) CreateProduct(ctx context.Context, product entity.Product) (*string, error) {
	g.logger.Debug("Inserting product into DB", "name", product.Name, "traceID", ctx.Value("traceID"))

	id := ulid.Make().String()
	_, err := g.db.Exec(`INSERT INTO products (product_id, name, description, price, quantity, created_at) VALUES ($1, $2, $3, $4, $5, 'NOW()');`,
		id,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity)
	if err != nil {
		g.logger.Error("Failed to insert product into db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return &id, nil
}
