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

func Test_CustomerGtw_CreateCustomer(t *testing.T) {
	type args struct {
		ctx      context.Context
		customer entity.Customer
	}

	aCustomerIncomplete := &entity.Customer{
		Name:    "John",
		Surname: "Doe",
	}

	scenarios := []struct {
		name        string
		args        args
		want        *string
		expectedErr error
	}{
		{"success", args{context.Background(), *aCustomer}, &customerID, nil},
		{"error", args{context.Background(), *aCustomerIncomplete}, nil, errors.New("error")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerGtw.On("CreateCustomer", tt.args.ctx, tt.args.customer).Return(tt.want, tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			got, err := customerGtw.CreateCustomer(tt.args.ctx, tt.args.customer)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
			customerGtw.AssertExpectations(t)
		})
	}
}

func Test_CustomerGtw_UpdateCustomer(t *testing.T) {
	type args struct {
		ctx      context.Context
		customer entity.Customer
	}

	errorID := "errorID"
	aCustomerNotFound := &entity.Customer{
		ID:      &errorID,
		Name:    "John",
		Surname: "Doe",
	}

	scenarios := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{"success", args{context.Background(), *aCustomer}, nil},
		{"error", args{context.Background(), *aCustomerNotFound}, errors.New("not found")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerGtw.On("UpdateCustomer", tt.args.ctx, tt.args.customer).Return(tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			err := customerGtw.UpdateCustomer(tt.args.ctx, tt.args.customer)

			assert.ErrorIs(t, err, tt.expectedErr)
			customerGtw.AssertExpectations(t)
		})
	}
}

func Test_CustomerGtw_DeleteCustomerByID(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerID string
	}

	scenarios := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{"success", args{context.Background(), customerID}, nil},
		{"error", args{context.Background(), "errorID"}, errors.New("not found")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerGtw.On("DeleteCustomerByID", tt.args.ctx, tt.args.customerID).Return(tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			err := customerGtw.DeleteCustomerByID(tt.args.ctx, tt.args.customerID)

			assert.ErrorIs(t, err, tt.expectedErr)
			customerGtw.AssertExpectations(t)
		})
	}
}
