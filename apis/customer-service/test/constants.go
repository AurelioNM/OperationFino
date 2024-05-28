package test

import (
	"cmd/customer-service/internal/domain/entity"
)

var CustomerID = "customerID"
var ErrorID = "errorID"

var ACustomer = &entity.Customer{
	ID:      &CustomerID,
	Name:    "John",
	Surname: "Doe",
	Email:   "john.doe@example.com",
}
var ACustomerArray = []*entity.Customer{
	{
		ID:      &CustomerID,
		Name:    "John",
		Surname: "Doe",
		Email:   "john.doe@mock.com",
	},
	{
		ID:      &CustomerID,
		Name:    "Aurelio",
		Surname: "Mororo",
		Email:   "aurelio@mock.com",
	},
}
