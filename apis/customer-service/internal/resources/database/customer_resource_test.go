package database

import (
	"cmd/customer-service/internal/domain/entity"
	"context"
	"log/slog"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

var (
	id = "01HWV4B6R3PG8XJW74P69FJTBH"

	customers = []*entity.Customer{
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
)

func TestGetCustomerList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error %s was not expected when mocking db conn", err)
	}
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	gateway := NewCustomerGateway(*logger, db)

	// ctx := context.WithValue(context.Background(), "traceID", "123456")

	rows := sqlmock.NewRows([]string{"customer_ID", "name", "surname", "email"}).
		AddRow("01HWV4B6R3PG8XJW74P69FJTBH", "John", "Doe", "john@gmail.com").
		AddRow("01HWV4B6R3PG8XJW74P69FJTBH", "Doe", "John", "doe@gmail.com")

	// Set up the expectations on the mock DB
	mock.ExpectQuery("SELECT customer_ID, name, surname, email FROM customers;").WillReturnRows(rows)

	// Call the function under test
	result, err := gateway.GetCustomerList(context.Background())

	// Check for errors
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Assert that the returned customers match the expected ones
	if !reflect.DeepEqual(result, customers) {
		t.Errorf("unexpected customers returned, got %v, want %v", result, customers)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
