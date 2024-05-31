package cache

import (
	"cmd/customer-service/internal/domain/entity"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type CustomerCache interface {
	GetCachedCustomer(ctx context.Context, customerID string) (*entity.Customer, error)
	CreateCustomerCache(ctx context.Context, customer entity.Customer) error
}

type customerCache struct {
	logger      slog.Logger
	client      *redis.Client
	customerKey string
}

func NewCustomerCache(l slog.Logger, client *redis.Client) CustomerCache {
	return &customerCache{
		logger:      *l.With("layer", "customer-cache"),
		client:      client,
		customerKey: "customer:",
	}
}

func (c *customerCache) GetCachedCustomer(ctx context.Context, customerID string) (*entity.Customer, error) {
	c.logger.Debug("Getting customer cache", "customerID", customerID, "traceID", ctx.Value("traceID"))
	key := c.customerKey + customerID

	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			c.logger.Error("Customer isn't cached", "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		return nil, err
	}
	c.logger.Debug("Got customer cache", "data", data, "traceID", ctx.Value("traceID"))

	var customer entity.Customer
	err = json.Unmarshal([]byte(data), &customer)
	if err != nil {
		c.logger.Error("Failed to unMarshal cached customer", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return &customer, nil
}

func (c *customerCache) CreateCustomerCache(ctx context.Context, customer entity.Customer) error {
	c.logger.Debug("Creating customer cache", "data", customer, "traceID", ctx.Value("traceID"))
	key := c.customerKey + *customer.ID

	data, err := json.Marshal(customer)
	if err != nil {
		c.logger.Error("Failed to marshal customer", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	err = c.client.Set(ctx, key, data, 0).Err() // 0 means no expiration
	if err != nil {
		c.logger.Error("Failed to create customer cache", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	return nil
}
