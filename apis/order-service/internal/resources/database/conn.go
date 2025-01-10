package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConn struct {
	DB *mongo.Client
}

func CreateDBConnPool(logger slog.Logger) (*DbConn, error) {
	logger.Debug("Creating MongoDB config")

	uri := os.Getenv("DB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Failed to create DB client: %s", err)
	}

	timeout, err := time.ParseDuration(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get DB_TIMEOUT from .env: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Failed to ping DB: %s", err)
	}

	logger.Debug("Conn pool created", "url", uri)

	return &DbConn{
		DB: client,
	}, nil
}
