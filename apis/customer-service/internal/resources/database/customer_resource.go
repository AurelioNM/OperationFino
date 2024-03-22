package database

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/internal/domain/gateway"

	"log/slog"
)

type customerGateway struct {
	logger    slog.Logger
	customers []*entity.Customer
}

func NewCustomerGateway(l slog.Logger) gateway.CustomerGateway {
	return &customerGateway{
		logger: *l.With("layer", "customer-gateway"),
		customers: []*entity.Customer{
			{ID: 1, Name: "Fulano", Surname: "Beltrano", Email: "fulano@gmail.com"},
			{ID: 2, Name: "Ciclano", Surname: "Nunes", Email: "ciclano@gmail.com"},
		},
	}
}

func (g *customerGateway) GetCustomerList() ([]*entity.Customer, error) {
	g.logger.Info("Getting all customers")
	return g.customers, nil
}
