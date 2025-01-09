package entity

import (
	"cmd/order-service/internal/resources/client/dto"
	"time"
)

type OrderRequest struct {
	CustomerEmail string                `json:"customer_email"`
	Products      []OrderRequestProduct `json:"products"`
}

type OrderRequestProduct struct {
	Name     string `json:"name" db:"name"`
	Quantity int64  `json:"quantity" db:"quantity"`
}

type Order struct {
	ID        *string       `json:"order_id" db:"order_id"`
	Customer  dto.Customer  `json:"customer" db:"customer"`
	Products  []dto.Product `json:"products" db:"products"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at" db:"updated_at"`
}
