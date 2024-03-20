package entity

type Customer struct {
	ID      int    `json:"customer_id" db:"customer_id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email   string `json:"email" db:"email"`
}
