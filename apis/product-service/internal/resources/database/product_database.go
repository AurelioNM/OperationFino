package database

import (
	"cmd/product-service/internal/domain/entity"
	"cmd/product-service/internal/domain/gateway"
	"cmd/product-service/internal/metrics"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/oklog/ulid/v2"
)

type productGateway struct {
	logger  slog.Logger
	metrics *metrics.ProductMetrics
	db      *sql.DB
}

func NewProductGateway(l slog.Logger, m *metrics.ProductMetrics, db *sql.DB) gateway.ProductGateway {
	return &productGateway{
		logger:  *l.With("layer", "product-database"),
		metrics: m,
		db:      db,
	}
}

func (g *productGateway) GetProductList(ctx context.Context) ([]*entity.Product, error) {
	g.logger.Debug("Getting all products from DB", "traceID", ctx.Value("traceID"))
	query := "SELECT product_id, name, description, price, quantity, created_at, updated_at FROM products;"
	start := time.Now()

	rows, err := g.db.Query(query)
	g.metrics.MeasureExternalDuration(start, "database", "ProductDB", "GetProductList", "")
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
	start := time.Now()

	rows, err := g.db.Query(query, productID)
	g.metrics.MeasureExternalDuration(start, "database", "ProductDB", "GetProductByID", "")
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

func (g *productGateway) GetProductByName(ctx context.Context, productName string) (*entity.Product, error) {
	g.logger.Debug("Getting product by name from db", "productName", productName, "traceID", ctx.Value("traceID"))
	query := "SELECT product_id, name, description, price, quantity, created_at, updated_at FROM products WHERE name = $1;"
	start := time.Now()

	rows, err := g.db.Query(query, productName)
	g.metrics.MeasureExternalDuration(start, "database", "ProductDB", "GetProductByName", "")
	if err != nil {
		g.logger.Error("Failed to get product by name from db", "error", err, "traceID", ctx.Value("traceID"))
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

	return nil, fmt.Errorf("No product found with name=%s", productName)
}

func (g *productGateway) CreateProduct(ctx context.Context, product entity.Product) (*string, error) {
	g.logger.Debug("Inserting product into DB", "name", product.Name, "traceID", ctx.Value("traceID"))
	start := time.Now()

	id := ulid.Make().String()
	_, err := g.db.Exec(`INSERT INTO products (product_id, name, description, price, quantity, created_at) VALUES ($1, $2, $3, $4, $5, 'NOW()');`,
		id,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity)
	g.metrics.MeasureExternalDuration(start, "database", "ProductDB", "CreateProduct", "")
	if err != nil {
		g.logger.Error("Failed to insert product into db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return &id, nil
}

func (g *productGateway) UpdateProduct(ctx context.Context, product entity.Product) error {
	g.logger.Debug("Updating product on db", "ID", product.ID, "traceID", ctx.Value("traceID"))
	start := time.Now()

	result, err := g.db.Exec(`UPDATE products SET name = $1, description = $2, price = $3, quantity = $4, updated_at = 'NOW()' WHERE product_id = $5;`,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.ID)
	g.metrics.MeasureExternalDuration(start, "database", "ProductDB", "UpdateProduct", "")
	if err != nil {
		g.logger.Error("Failed to update product on db", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return validateIfRowWasAffected(result, *product.ID)
}

func (g *productGateway) DeleteProductByID(ctx context.Context, productID string) error {
	g.logger.Debug("Deleting product on db", "ID", productID, "traceID", ctx.Value("traceID"))
	start := time.Now()

	result, err := g.db.Exec(`DELETE FROM products WHERE product_id = $1;`, productID)
	g.metrics.MeasureExternalDuration(start, "database", "ProductDB", "DeleteProductByID", "")
	if err != nil {
		g.logger.Error("Failed to update product on db", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return validateIfRowWasAffected(result, productID)
}

func validateIfRowWasAffected(result sql.Result, productID string) error {
	rows, err := result.RowsAffected()
	if rows == 0 || err != nil {
		return fmt.Errorf("Product not found with ID=%s", productID)
	}

	return nil
}
