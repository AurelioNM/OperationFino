package main

import (
	"cmd/customer-service/internal/api"
	"cmd/customer-service/internal/domain/service"
	"cmd/customer-service/internal/metrics"
	"cmd/customer-service/internal/pyroscope"
	"cmd/customer-service/internal/resources/cache"
	"cmd/customer-service/internal/resources/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"log/slog"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Profiler
	if pyroscope.StartPyroscope() {
		defer pyroscope.WaitPyroscope()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file", "error", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	db, err := database.CreateDBConnPool(*logger)
	if err != nil {
		logger.Error("Error creating DB", "error", err)
		return
	}

	cacheClient, err := cache.GetCacheClient(*logger)
	if err != nil {
		logger.Error("Error creating cache client", "error", err)
		return
	}

	// Metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	prometheusHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	metrics := metrics.NewCustomerMetrics(*logger, reg)
	customerGtw := database.NewCustomerGateway(*logger, db.DB)
	customerCache := cache.NewCustomerCache(*logger, cacheClient)
	customerSvc := service.NewCustomerService(*logger, customerGtw, customerCache)
	customerHandler := api.NewCustomerHandler(*logger, metrics, customerSvc)

	r := createRouter(prometheusHandler, customerHandler)
	logger.Debug("Starting customer-service", "port", os.Getenv("APP_PORT"))
	go http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r)

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	logger.Debug("Stoping customer-service")
}

func createRouter(prometheusHandler http.Handler, customerHandler api.CustomerHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/metrics", prometheusHandler.ServeHTTP).Methods("GET")
	r.HandleFunc("/v1/customers", customerHandler.GetCustomers).Methods("GET")
	r.HandleFunc("/v1/customers/{id}", customerHandler.GetCustomerByID).Methods("GET")
	r.HandleFunc("/v2/customers/{id}", customerHandler.V2GetCustomerByID).Methods("GET")
	r.HandleFunc("/v1/customers/email/{email}", customerHandler.GetCustomerByEmail).Methods("GET")
	r.HandleFunc("/v2/customers/email/{email}", customerHandler.V2GetCustomerByEmail).Methods("GET")
	r.HandleFunc("/v1/customers/name/{name}", customerHandler.GetCustomerByName).Methods("GET")
	r.HandleFunc("/v2/customers/name/{name}", customerHandler.V2GetCustomerByName).Methods("GET")
	r.HandleFunc("/v1/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/v1/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/v1/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	// Swagger
	r.PathPrefix("/customers/doc/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger.yml"),
	))
	fs := http.FileServer(http.Dir("./"))
	r.PathPrefix("/").Handler(fs)

	return r
}
