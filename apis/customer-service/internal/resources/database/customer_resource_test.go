package database

import (
	"cmd/customer-service/internal/domain/entity"
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Select(dest interface{}, query string, args ...interface{}) error {
	argsMock := m.Called(dest, query, args)
	return argsMock.Error(0)
}

func TestGetCustomerList(t *testing.T) {
	mockDB := new(MockDB)
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	gateway := NewCustomerGateway(&logger, mockDB)

	id := "01HWV4B6R3PG8XJW74P69FJTBH"

	customers := []*entity.Customer{
		{
			ID:        &id,
			Name:      "John",
			Surname:   "Doe",
			Email:     "john@gmail.com",
			Birthdate: entity.DateOnly{Time: time.Now()},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        &id,
			Name:      "Doe",
			Surname:   "John",
			Email:     "doe@gmail.com",
			Birthdate: entity.DateOnly{Time: time.Now()},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockDB.On("Select", &customers, "SELECT customer_ID, name, surname, email FROM customers;").Return(nil)

	ctx := context.WithValue(context.Background(), "traceID", "123456")
	result, err := gateway.GetCustomerList(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, customers, result)

	mockDB.AssertExpectations(t)
}

func TestGetCustomerList2(t *testing.T) {
	// Create a mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a logger mock
	loggerMock := &slog.MockLogger{}

	// Set up mock expectations
	rows := sqlmock.NewRows([]string{"customer_id", "name", "surname", "email"}).
		AddRow("1", "John", "Doe", "john@example.com").
		AddRow("2", "Jane", "Smith", "jane@example.com")
	mock.ExpectQuery("SELECT customer_ID, name, surname, email FROM customers;").WillReturnRows(rows)

	// Create the customer gateway with mocks
	gateway := NewCustomerGateway(loggerMock, db)

	// Call the method
	ctx := context.WithValue(context.Background(), "traceID", "test-trace-id")
	customers, err := gateway.GetCustomerList(ctx)

	// Check expectations
	assert.NoError(t, err)
	assert.Len(t, customers, 2)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
