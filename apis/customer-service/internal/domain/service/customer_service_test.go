package service

import (
	"cmd/customer-service/internal/domain/entity"
	"cmd/customer-service/mocks"
	"cmd/customer-service/test"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var customerSvc = new(mocks.CustomerService)

func Test_CustomerSvc_GetCustomerList(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	scenarios := []struct {
		name        string
		args        args
		want        []*entity.Customer
		expectedErr error
	}{
		{"success", args{context.Background()}, test.ACustomerArray, nil},
		{"error", args{context.TODO()}, nil, errors.New("error")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerSvc.On("GetCustomerList", tt.args.ctx).Return(tt.want, tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			got, err := customerSvc.GetCustomerList(tt.args.ctx)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
			customerSvc.AssertExpectations(t)
		})
	}
}

func Test_CustomerSvc_GetCustomerByID(t *testing.T) {
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
		{"success", args{context.Background(), test.CustomerID}, test.ACustomer, nil},
		{"error", args{context.Background(), "errorID"}, nil, errors.New("error")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerSvc.On("GetCustomerByID", tt.args.ctx, tt.args.customerID).Return(tt.want, tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			got, err := customerSvc.GetCustomerByID(tt.args.ctx, tt.args.customerID)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
			customerSvc.AssertExpectations(t)
		})
	}
}

func Test_CustomerSvc_CreateCustomer(t *testing.T) {
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
		{"success", args{context.Background(), *test.ACustomer}, &test.CustomerID, nil},
		{"error", args{context.Background(), *aCustomerIncomplete}, nil, errors.New("error")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerSvc.On("CreateCustomer", tt.args.ctx, tt.args.customer).Return(tt.want, tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			got, err := customerSvc.CreateCustomer(tt.args.ctx, tt.args.customer)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
			customerSvc.AssertExpectations(t)
		})
	}
}

func Test_CustomerSvc_UpdateCustomer(t *testing.T) {
	type args struct {
		ctx      context.Context
		customer entity.Customer
	}

	aCustomerNotFound := &entity.Customer{
		ID:      &test.ErrorID,
		Name:    "John",
		Surname: "Doe",
	}

	scenarios := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{"success", args{context.Background(), *test.ACustomer}, nil},
		{"error", args{context.Background(), *aCustomerNotFound}, errors.New("not found")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerSvc.On("UpdateCustomer", tt.args.ctx, tt.args.customer).Return(tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			err := customerSvc.UpdateCustomer(tt.args.ctx, tt.args.customer)

			assert.ErrorIs(t, err, tt.expectedErr)
			customerSvc.AssertExpectations(t)
		})
	}
}

func Test_CustomerSvc_DeleteCustomerByID(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerID string
	}

	scenarios := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{"success", args{context.Background(), test.CustomerID}, nil},
		{"error", args{context.Background(), "errorID"}, errors.New("not found")},
	}

	for _, tt := range scenarios {
		tt := tt

		customerSvc.On("DeleteCustomerByID", tt.args.ctx, tt.args.customerID).Return(tt.expectedErr)

		t.Run(tt.name, func(t *testing.T) {
			err := customerSvc.DeleteCustomerByID(tt.args.ctx, tt.args.customerID)

			assert.ErrorIs(t, err, tt.expectedErr)
			customerSvc.AssertExpectations(t)
		})
	}
}
