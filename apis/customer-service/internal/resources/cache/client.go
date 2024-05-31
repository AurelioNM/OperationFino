package cache

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetCacheClient(logger slog.Logger) (*redis.Client, error) {
	logger.Debug("Creating cache client")
	connTimeout, err := time.ParseDuration(os.Getenv("CACHE_CONN_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get CACHE_CONN_TIMEOUT from .env: %s", err)
	}

	readTimeout, err := time.ParseDuration(os.Getenv("CACHE_READ_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get CACHE_READ_TIMEOUT from .env: %s", err)
	}

	writeTimeout, err := time.ParseDuration(os.Getenv("CACHE_WRITE_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get CACHE_WRITE_TIMEOUT from .env: %s", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("CACHE_ADDRESS"),
		Password:     "",
		DB:           0,
		DialTimeout:  connTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping DB: %s", err)
	}

	logger.Debug("Conn pool created",
		slog.String("pong", pong),
		slog.String("ConnTimeout", connTimeout.String()),
		slog.String("ReadTimeout", readTimeout.String()),
		slog.String("WriteTimeout", writeTimeout.String()))

	return client, nil
}
