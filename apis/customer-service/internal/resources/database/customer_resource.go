package database

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"
	"context"
	"database/sql"
	"fmt"

	"log/slog"

	"github.com/oklog/ulid/v2"
)

type customerGateway struct {
	logger slog.Logger
	db     *sql.DB
}

func NewCustomerGateway(l slog.Logger, db *sql.DB) gateway.CustomerGateway {
	return &customerGateway{
		logger: *l.With("layer", "customer-gateway"),
		db:     db,
	}
}

func (g *customerGateway) GetCustomerList(ctx context.Context) ([]*entity.Customer, error) {
	g.logger.Debug("Getting all customers from db", "traceID", ctx.Value("traceID"))
	query := "SELECT customer_ID, name, surname, email, birthdate, created_at, updated_at FROM customers;"

	rows, err := g.db.Query(query)
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
			g.logger.Error("Error scaning row", "error", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (g *customerGateway) GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error) {
	g.logger.Debug("Getting customer by ID from db", "ID", customerID, "traceID", ctx.Value("traceID"))
	query := "SELECT customer_id, name, surname, email, birthdate, created_at, updated_at FROM customers WHERE customer_id = $1;"

	rows, err := g.db.Query(query, customerID)
	if err != nil {
		g.logger.Error("Failed to get customer by ID from db", "error", err, "traceID", ctx.Value("traceID"))
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

	return nil, fmt.Errorf("No customer found ID=%s", customerID)
}

func (g *customerGateway) GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error) {
	g.logger.Debug("Getting customer by email from db", "email", customerEmail, "traceID", ctx.Value("traceID"))
	query := "SELECT customer_id, name, surname, email, birthdate, created_at, updated_at FROM customers WHERE email = $1;"

	rows, err := g.db.Query(query, customerEmail)
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

	return nil, fmt.Errorf("No customer found email=%s", customerEmail)
}

func (g *customerGateway) CreateCustomer(ctx context.Context, customer entity.Customer) (*string, error) {
	g.logger.Debug("Inserting customer into db", "email", customer.Email, "traceID", ctx.Value("traceID"))

	id := ulid.Make().String()
	_, err := g.db.Exec(`INSERT INTO customers (customer_id, name, surname, email, birthdate, created_at) VALUES ($1, $2, $3, $4, $5, 'NOW()');`,
		id,
		customer.Name,
		customer.Surname,
		customer.Email,
		customer.Birthdate)
	if err != nil {
		g.logger.Error("Failed to insert customer into db", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return &id, nil
}

func (g *customerGateway) UpdateCustomer(ctx context.Context, customer entity.Customer) error {
	g.logger.Debug("Updating customer on db", "ID", customer.ID, "traceID", ctx.Value("traceID"))

	result, err := g.db.Exec(`UPDATE customers SET name = $1, surname = $2, email = $3, updated_at = 'NOW()' WHERE customer_id = $4;`,
		customer.Name,
		customer.Surname,
		customer.Email,
		customer.ID)
	if err != nil {
		g.logger.Error("Failed to update customer on db", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return validateIfCustomerExists(result, *customer.ID)
}

func (g *customerGateway) DeleteCustomerByID(ctx context.Context, customerID string) error {
	g.logger.Debug("Deleting customer on db", "ID", customerID, "traceID", ctx.Value("traceID"))

	result, err := g.db.Exec(`DELETE FROM customers WHERE customer_id = $1;`, customerID)
	if err != nil {
		g.logger.Error("Failed to update customer on db", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return validateIfCustomerExists(result, customerID)
}

func validateIfCustomerExists(result sql.Result, customerID string) error {
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("Customer not found ID=%s", customerID)
	}

	return nil
}
