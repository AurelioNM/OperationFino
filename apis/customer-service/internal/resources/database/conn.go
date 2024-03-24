package database

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"log/slog"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DbConnPool struct {
	DB *sqlx.DB
}

func CreateDBConnPool(logger slog.Logger) (*DbConnPool, error) {
	logger.Debug("Creating DSN config")
	url := &url.URL{
		Scheme: os.Getenv("DB_TYPE"),
		Host: fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User: url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Path: os.Getenv("DB_NAME"),
	}

	timeout, err := time.ParseDuration(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get DB_TIMEOUT from .env: %s", err)
	}
	connOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get DB_MAX_OPEN_CONN from .env: %s", err)
	}

	logger.Debug("Creating DB conn pool")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	
	scheme := "pgx"
	connPool, err := sqlx.ConnectContext(ctx, scheme, url.String())
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize DB conn pool: %s", err)
	}

	connPool.SetConnMaxLifetime(timeout)
	connPool.SetMaxIdleConns(connOpen)
	connPool.SetMaxOpenConns(connOpen)

	logger.Debug("Conn pool created", 
		slog.String("url", url.String()),
		slog.String("ConnTimeout", timeout.String()),
		slog.Int("MaxConnLifetime", connOpen),
		slog.Int("SetMaxOpenConns", connOpen))

	return &DbConnPool{
		DB: connPool,
	}, nil
}
