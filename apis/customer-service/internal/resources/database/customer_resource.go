package database

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"

	"log/slog"

	"github.com/jmoiron/sqlx"
)

type customerGateway struct {
	logger    slog.Logger
	customers []*entity.Customer
	db        *sqlx.DB
}

func NewCustomerGateway(l slog.Logger, db *sqlx.DB) gateway.CustomerGateway {
	return &customerGateway{
		logger: *l.With("layer", "customer-gateway"),
		customers: []*entity.Customer{
			{ID: 1, Name: "Fulano", Surname: "Beltrano", Email: "fulano@gmail.com"},
			{ID: 2, Name: "Ciclano", Surname: "Nunes", Email: "ciclano@gmail.com"},
		},
		db: db,
	}
}

func (g *customerGateway) GetCustomerList() ([]*entity.Customer, error) {
	customers := []entity.Customer{}
	err := g.db.Select(&customers, "SELECT customer_ID, name, surname, email FROM customers")
	if err != nil {
		g.logger.Error("Failed to get customers from db")
	}

	g.logger.Info("Getting all customers")
	return g.customers, nil
}

