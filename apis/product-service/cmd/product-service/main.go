package main

import (
	"cmd/product-service/internal/api"
	"cmd/product-service/internal/domain/service"
	"cmd/product-service/internal/resources/database"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// TODO pyroscope config

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

	// TODO metrics config

	productGtw := database.NewProductGateway(*logger, db.DB)
	productSvc := service.NewProductService(*logger, productGtw)
	productHandler := api.NewProductHandler(*logger, productSvc)

	r := createRouter(productHandler)
	logger.Debug("Starting prodduct-service", "port", os.Getenv("APP_PORT"))
	go http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r)

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	logger.Debug("Stoping product-service")
}

func createRouter(productHandler api.ProductHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/v1/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/v1/products", productHandler.CreateProduct).Methods("POST")

	r.PathPrefix("/products/doc/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger.yml"),
	))
	fs := http.FileServer(http.Dir("./"))
	r.PathPrefix("/").Handler(fs)

	return r
}
