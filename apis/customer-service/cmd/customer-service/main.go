package main

import (
	"cmd/customer-service/internal/api"
	"log"
	"net/http"
	"os"

	"log/slog"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	customerHandler := api.NewCustomerHandler(logger)

	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	logger.Debug("Running customer-service on port 8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
