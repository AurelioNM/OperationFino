package database

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCustomerByID(t *testing.T) {
	// Create a new instance of the mock CustomerGateway
	mockCustomerGateway := new(mocks.CustomerGateway)

	// Create a sample customer
	customerID := "customerID"
	expectedCustomer := &entity.Customer{
		ID:      &customerID,
		Name:    "John",
		Surname: "Doe",
		Email:   "john.doe@example.com",
	}

	// Set up the expectation for GetCustomerByID
	mockCustomerGateway.On("GetCustomerByID", mock.Anything, customerID).Return(expectedCustomer, nil)

	// Call the method
	ctx := context.TODO()
	customer, err := mockCustomerGateway.GetCustomerByID(ctx, customerID)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	// Ensure that the expectations were met
	mockCustomerGateway.AssertExpectations(t)
}
