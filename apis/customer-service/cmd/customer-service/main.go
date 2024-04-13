package main

import (
	"cmd/customer-service/internal/api"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/metrics"
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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	db, err := database.CreateDBConnPool(*logger)
	if err != nil {
		logger.Error("Error creating DB", "error", err)
		return
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	metrics := metrics.NewCustomerMetrics(*logger, reg)
	customerGtw := database.NewCustomerGateway(*logger, db.DB)
	customerSvc := service.NewCustomerService(*logger, customerGtw)
	customerHandler := api.NewCustomerHandler(*logger, metrics, customerSvc)

	r.HandleFunc("/metrics", promHandler.ServeHTTP).Methods("GET")
	r.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", customerHandler.GetCustomerByID).Methods("GET")
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	logger.Debug("Running customer-service", "port", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r))
}
