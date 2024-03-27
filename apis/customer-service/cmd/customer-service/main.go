package main

import (
	"cmd/customer-service/internal/api"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/pyroscope"
	"cmd/customer-service/internal/resources/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"log/slog"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if pyroscope.StartPyroscope() {
		defer pyroscope.WaitPyroscope()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file", "error", err)
	}

	r := mux.NewRouter()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	r.Use(loggingMiddleware(logger))

	db, err := database.CreateDBConnPool(*logger)
	if err != nil {
		logger.Error("Error creating DB", "error", err)
		return
	}

	customerGtw := database.NewCustomerGateway(*logger, db.DB)
	customerSvc := service.NewCustomerService(*logger, customerGtw)
	customerHandler := api.NewCustomerHandler(*logger, customerSvc)

	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", customerHandler.GetCustomerByID).Methods("GET")
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	logger.Debug("Running customer-service", "port2", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r))
}

func loggingMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debug(fmt.Sprintf("Recieved request %s %s", r.Method, r.URL.Path))
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
