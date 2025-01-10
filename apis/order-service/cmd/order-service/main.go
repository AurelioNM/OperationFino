package main

import (
	"cmd/order-service/internal/api"
	"cmd/order-service/internal/domain/service"
	"cmd/order-service/internal/metrics"
	"cmd/order-service/internal/pyroscope"
	"cmd/order-service/internal/resources/client"
	"cmd/order-service/internal/resources/database"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
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

	// Metrics
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	prometheusHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	metrics := metrics.NewOrderMetrics(*logger, reg)

	customerGtw := client.NewCustomerGateway(*logger, metrics)
	productGtw := client.NewProductGateway(*logger, metrics)
	orderGtw := database.NewOrderGateway(*logger, db.DB)
	orderSvc := service.NewOrderService(*logger, orderGtw, customerGtw, productGtw)
	orderHandler := api.NewOrderHandler(*logger, metrics, orderSvc)

	r := createRouter(prometheusHandler, orderHandler)
	logger.Debug("Starting prodduct-service", "port", os.Getenv("APP_PORT"))
	go http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r)

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	logger.Debug("Stoping order-service")
}

func createRouter(prometheusHandler http.Handler, orderHandler api.OrderHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/metrics", prometheusHandler.ServeHTTP).Methods("GET")
	r.HandleFunc("/v1/orders/{id}", orderHandler.GetOrderByID).Methods("GET")
	r.HandleFunc("/v1/orders/{id}", orderHandler.DeleteOrderByID).Methods("DELETE")
	r.HandleFunc("/v1/orders/customers/{customerID}", orderHandler.GetOrdersByCustomerID).Methods("GET")
	r.HandleFunc("/v1/orders", orderHandler.CreateOrder).Methods("POST")

	r.PathPrefix("/orders/doc/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger.yml"),
	))
	fs := http.FileServer(http.Dir("./"))
	r.PathPrefix("/").Handler(fs)

	return r
}
