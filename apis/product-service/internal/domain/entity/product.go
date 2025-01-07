package entity

import "time"

type Product struct {
	ID          *string `json:"product_id" db:"product_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int64   `json:"quantity" db:"quantity"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
