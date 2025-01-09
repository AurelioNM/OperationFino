package dto

type GetCustomerByEmailResponseDTO struct {
	Message     string      `json:"message"`
	Timestamp   string      `json:"timestamp"`
	ElapsedTime string      `json:"elapsed_time"`
	Data        CustomerDTO `json:"data"`
}

type CustomerDTO struct {
	Customer Customer `json:"customer"`
}

type Customer struct {
	ID      *string `json:"customer_id" db:"customer_id"`
	Name    string  `json:"name" db:"name"`
	Surname string  `json:"surname" db:"surname"`
	Email   string  `json:"email" db:"email"`
}
