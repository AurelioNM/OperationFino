// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	entity "cmd/customer-service/internal/domain/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CustomerGateway is an autogenerated mock type for the CustomerGateway type
type CustomerGateway struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: ctx, customer
func (_m *CustomerGateway) CreateCustomer(ctx context.Context, customer entity.Customer) (*string, error) {
	ret := _m.Called(ctx, customer)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomer")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Customer) (*string, error)); ok {
		return rf(ctx, customer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Customer) *string); ok {
		r0 = rf(ctx, customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Customer) error); ok {
		r1 = rf(ctx, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomerByID provides a mock function with given fields: ctx, customerID
func (_m *CustomerGateway) DeleteCustomerByID(ctx context.Context, customerID string) error {
	ret := _m.Called(ctx, customerID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCustomerByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, customerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCustomerByEmail provides a mock function with given fields: ctx, customerEmail
func (_m *CustomerGateway) GetCustomerByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error) {
	ret := _m.Called(ctx, customerEmail)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomerByEmail")
	}

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Customer, error)); ok {
		return rf(ctx, customerEmail)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Customer); ok {
		r0 = rf(ctx, customerEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByID provides a mock function with given fields: ctx, customerID
func (_m *CustomerGateway) GetCustomerByID(ctx context.Context, customerID string) (*entity.Customer, error) {
	ret := _m.Called(ctx, customerID)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomerByID")
	}

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Customer, error)); ok {
		return rf(ctx, customerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Customer); ok {
		r0 = rf(ctx, customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByName provides a mock function with given fields: ctx, customerName
func (_m *CustomerGateway) GetCustomerByName(ctx context.Context, customerName string) (*entity.Customer, error) {
	ret := _m.Called(ctx, customerName)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomerByName")
	}

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Customer, error)); ok {
		return rf(ctx, customerName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Customer); ok {
		r0 = rf(ctx, customerName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerList provides a mock function with given fields: ctx
func (_m *CustomerGateway) GetCustomerList(ctx context.Context) ([]*entity.Customer, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomerList")
	}

	var r0 []*entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.Customer, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Customer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCustomer provides a mock function with given fields: ctx, customer
func (_m *CustomerGateway) UpdateCustomer(ctx context.Context, customer entity.Customer) error {
	ret := _m.Called(ctx, customer)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCustomer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Customer) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCustomerGateway creates a new instance of CustomerGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCustomerGateway(t interface {
	mock.TestingT
	Cleanup(func())
}) *CustomerGateway {
	mock := &CustomerGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
