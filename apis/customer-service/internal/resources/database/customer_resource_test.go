package database

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var customerGtw = new(mocks.CustomerGateway)
var customerID = "customerID"
var aCustomer = &entity.Customer{
	ID:      &customerID,
	Name:    "John",
	Surname: "Doe",
	Email:   "john.doe@example.com",
}
var aCustomerArray = []*entity.Customer{
	{
		ID:      &customerID,
		Name:    "John",
		Surname: "Doe",
		Email:   "john.doe@mock.com",
	},
	{
		ID:      &customerID,
		Name:    "Aurelio",
		Surname: "Mororo",
		Email:   "aurelio@mock.com",
	},
}

func Test_CustomerGtw_GetCustomerByID(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerID string
	}

	scenarios := []struct {
		name        string
		args        args
		want        *entity.Customer
		expectedErr error
	}{
		{"success", args{context.Background(), customerID}, aCustomer, nil},
		{"error", args{context.Background(), "errorID"}, nil, errors.New("error")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerGtw.On("GetCustomerByID", tt.args.ctx, tt.args.customerID).Return(tt.want, tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			got, err := customerGtw.GetCustomerByID(tt.args.ctx, tt.args.customerID)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
			customerGtw.AssertExpectations(t)
		})
	}
}

func Test_CustomerGtw_GetCustomerList(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	scenarios := []struct {
		name        string
		args        args
		want        []*entity.Customer
		expectedErr error
	}{
		{"success", args{context.Background()}, aCustomerArray, nil},
		{"error", args{context.TODO()}, nil, errors.New("error")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerGtw.On("GetCustomerList", tt.args.ctx).Return(tt.want, tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			got, err := customerGtw.GetCustomerList(tt.args.ctx)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
			customerGtw.AssertExpectations(t)
		})
	}
}
