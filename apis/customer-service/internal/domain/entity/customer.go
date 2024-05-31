package entity

import (
	"time"
)

type Customer struct {
	ID        *string    `json:"customer_id" db:"customer_id"`
	Name      string     `json:"name" db:"name"`
	Surname   string     `json:"surname" db:"surname"`
	Email     string     `json:"email" db:"email"`
	Birthdate string     `json:"birthdate" db:"birthdate"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
