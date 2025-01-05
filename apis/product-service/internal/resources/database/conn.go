package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DbConn struct {
	DB *sql.DB
}

func CreateDBConnPool(logger slog.Logger) (*DbConn, error) {
	logger.Debug("Creating DB DSN config")
	url := &url.URL{
		Scheme: os.Getenv("DB_TYPE"),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:   url.UserPassword(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")),
		Path:   os.Getenv("DB_NAME"),
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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db, err := sql.Open("pgx", url.String())
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize DB conn pool: %s", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Failed to ping DB: %s", err)
	}

	db.SetConnMaxLifetime(timeout)
	db.SetMaxIdleConns(connOpen)
	db.SetMaxOpenConns(connOpen)

	logger.Debug("Conn pool created",
		slog.String("url", url.String()),
		slog.String("ConnTimeout", timeout.String()),
		slog.Int("MaxConnLifetime", connOpen),
		slog.Int("SetMaxOpenConn", connOpen))

	return &DbConn{
		DB: db,
	}, nil
}
