package entity

import "time"

type OrderRequest struct {
	CustomerEmail string                `json:"name"`
	Products      []OrderRequestProduct `json:"description"`
}

type OrderRequestProduct struct {
	Name     string `json:"name" db:"name"`
	Quantity int64  `json:"quantity" db:"quantity"`
}

type Order struct {
	ID        *string    `json:"order_id" db:"order_id"`
	Customer  Customer   `json:"customer" db:"customer"`
	Products  []Product  `json:"products" db:"products"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type Product struct {
	ID          *string `json:"order_id" db:"order_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int64   `json:"quantity" db:"quantity"`
}

type Customer struct {
	ID    *string `json:"customer_id" db:"customer_id"`
	Name  string  `json:"name" db:"name"`
	Email string  `json:"email" db:"email"`
}
