package database

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"
	"cmd/customer-service/internal/metrics"
	"context"
	"database/sql"
	"fmt"
	"time"

	"log/slog"

	"github.com/oklog/ulid/v2"
)

type customerGateway struct {
	logger  slog.Logger
	metrics *metrics.CustomerMetrics
	db      *sql.DB
}

func NewCustomerGateway(l slog.Logger, m *metrics.CustomerMetrics, db *sql.DB) gateway.CustomerGateway {
	return &customerGateway{
		logger:  *l.With("layer", "customer-database"),
		metrics: m,
		db:      db,
	}
}

func (g *customerGateway) GetCustomerList(ctx context.Context) ([]*entity.Customer, error) {
	g.logger.Debug("Getting all customers from db", "traceID", ctx.Value("traceID"))
	query := "SELECT customer_ID, name, surname, email, birthdate, created_at, updated_at FROM customers;"
	start := time.Now()

	rows, err := g.db.Query(query)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "GetCustomerList", "")
	if err != nil {
		g.logger.Error("Failed to get customers from db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer rows.Close()
	customers := make([]*entity.Customer, 0)
	for rows.Next() {
		customer := &entity.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.Email, &customer.Birthdate, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			g.logger.Error("Error scaning row", "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		customers = append(customers, customer)
	}

	g.logger.Info("Found customer list on DB", "size", len(customers))
	return customers, nil
}

func (g *customerGateway) GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error) {
	g.logger.Debug("Getting customer by ID from db", "ID", customerID, "traceID", ctx.Value("traceID"))
	query := "SELECT customer_id, name, surname, email, birthdate, created_at, updated_at FROM customers WHERE customer_id = $1;"
	start := time.Now()

	rows, err := g.db.Query(query, customerID)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "GetCustomerByID", "")
	if err != nil {
		g.logger.Error("Failed to get customer by ID from db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		customer := entity.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.Email, &customer.Birthdate, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			g.logger.Error("Error scaning product row", "error", err)
			return nil, err
		}
		return &customer, nil
	}

	return nil, fmt.Errorf("No customer found with ID=%s", customerID)
}

func (g *customerGateway) GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error) {
	g.logger.Debug("Getting customer by email from db", "email", customerEmail, "traceID", ctx.Value("traceID"))
	query := "SELECT customer_id, name, surname, email, birthdate, created_at, updated_at FROM customers WHERE email = $1;"
	start := time.Now()

	rows, err := g.db.Query(query, customerEmail)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "GetCustomerByEmail", "")
	if err != nil {
		g.logger.Error("Failed to get customer by email from db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		customer := entity.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.Email, &customer.Birthdate, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			g.logger.Error("Error scaning row", "error", err)
			return nil, err
		}
		return &customer, nil
	}

	return nil, fmt.Errorf("No customer found with email=%s", customerEmail)
}

func (g *customerGateway) GetCustomerByName(ctx context.Context, customerName string) (*entity.Customer, error) {
	g.logger.Debug("Getting customer by name from db", "name", customerName, "traceID", ctx.Value("traceID"))
	query := "SELECT customer_id, name, surname, email, birthdate, created_at, updated_at FROM customers WHERE name = $1 LIMIT 1;"
	start := time.Now()

	rows, err := g.db.Query(query, customerName)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "GetCustomerByName", "")
	if err != nil {
		g.logger.Error("Failed to get customer by name from db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		customer := entity.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.Email, &customer.Birthdate, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			g.logger.Error("Error scaning row", "error", err)
			return nil, err
		}
		return &customer, nil
	}

	return nil, fmt.Errorf("No customer found with name=%s", customerName)
}

func (g *customerGateway) CreateCustomer(ctx context.Context, customer entity.Customer) (*string, error) {
	g.logger.Debug("Inserting customer into db", "email", customer.Email, "traceID", ctx.Value("traceID"))
	start := time.Now()

	id := ulid.Make().String()
	_, err := g.db.Exec(`INSERT INTO customers (customer_id, name, surname, email, birthdate, created_at) VALUES ($1, $2, $3, $4, $5, 'NOW()');`,
		id,
		customer.Name,
		customer.Surname,
		customer.Email,
		customer.Birthdate)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "CreateCustomer", "")
	if err != nil {
		g.logger.Error("Failed to insert customer into db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return &id, nil
}

func (g *customerGateway) UpdateCustomer(ctx context.Context, customer entity.Customer) error {
	g.logger.Debug("Updating customer on db", "ID", customer.ID, "traceID", ctx.Value("traceID"))
	start := time.Now()

	result, err := g.db.Exec(`UPDATE customers SET name = $1, surname = $2, email = $3, updated_at = 'NOW()' WHERE customer_id = $4;`,
		customer.Name,
		customer.Surname,
		customer.Email,
		customer.ID)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "UpdateCustomer", "")
	if err != nil {
		g.logger.Error("Failed to update customer on db", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return validateIfRowWasAffected(result, *customer.ID)
}

func (g *customerGateway) DeleteCustomerByID(ctx context.Context, customerID string) error {
	g.logger.Debug("Deleting customer on db", "ID", customerID, "traceID", ctx.Value("traceID"))
	start := time.Now()

	result, err := g.db.Exec(`DELETE FROM customers WHERE customer_id = $1;`, customerID)
	g.metrics.MeasureExternalDuration(start, "database", "CustomerDB", "DeleteCustomerByID", "")
	if err != nil {
		g.logger.Error("Failed to update customer on db", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return validateIfRowWasAffected(result, customerID)
}

func validateIfRowWasAffected(result sql.Result, customerID string) error {
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("Customer not found ID=%s", customerID)
	}

	return nil
}
