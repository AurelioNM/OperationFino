package dto

type GetProductByNameResponseDTO struct {
	Message     string     `json:"message"`
	Timestamp   string     `json:"timestamp"`
	ElapsedTime string     `json:"elapsed_time"`
	Data        ProductDTO `json:"data"`
}

type ProductDTO struct {
	Product Product `json:"product"`
}

type Product struct {
	ID          *string `json:"product_id" db:"product_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int64   `json:"quantity" db:"quantity"`
}
