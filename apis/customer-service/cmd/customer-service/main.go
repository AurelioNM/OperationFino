package main

import (
	"cmd/customer-service/internal/api"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/resources/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"log/slog"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DbConnPool struct {
	DB *sqlx.DB
}

func createDBConnPool(logger slog.Logger) (*DbConnPool, error) {
	logger.Debug("Creating DSN config")
	url := &url.URL{
		Scheme: "postgres",
		Host: fmt.Sprintf("%s:%d", "localhost", 5441),
		User: url.UserPassword("customer", "customer"),
		Path: "customer-service",
	}

	logger.Debug("Creating DB conn pool")

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 999999)
	// defer cancel()
	
	scheme := "pgx"
	connPool, err := sqlx.ConnectContext(ctx, scheme, url.String())
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize DB conn pool: %s", err)
	}

	connPool.SetConnMaxLifetime(5000)
	connPool.SetMaxIdleConns(10)
	connPool.SetMaxOpenConns(10)

	logger.Debug("Conn pool created", 
		slog.String("url", url.String()),
		slog.Int("ConnTimeout", 5000),
		slog.Int("MaxConnLifetime", 10),
		slog.Int("SetMaxOpenConns", 10),
	)

	return &DbConnPool{
		DB: connPool,
	}, nil
}

func main() {
	r := mux.NewRouter()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	db, err := createDBConnPool(*logger)
	if err != nil {
		logger.Error("Error creating DB", "error", err)
		return
	}

	customerGtw := database.NewCustomerGateway(*logger, db.DB)
	customerSvc := service.NewCustomerService(*logger, customerGtw)
	customerHandler := api.NewCustomerHandler(*logger, customerSvc)

	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	logger.Debug("Running customer-service on port 8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
