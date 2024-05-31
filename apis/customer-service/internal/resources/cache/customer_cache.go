package cache

import (
	"cmd/customer-service/internal/domain/entity"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type CustomerCache interface {
	ReadCacheByID(ctx context.Context, customerID string) (*entity.Customer, error)
	WriteCacheByID(ctx context.Context, customer entity.Customer) error
	ReadCacheByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error)
	WriteCacheByEmail(ctx context.Context, customerEmail entity.Customer) error
}

type customerCache struct {
	logger   slog.Logger
	client   *redis.Client
	idKey    string
	emailKey string
}

func NewCustomerCache(l slog.Logger, client *redis.Client) CustomerCache {
	return &customerCache{
		logger:   *l.With("layer", "customer-cache"),
		client:   client,
		idKey:    "customer-id:",
		emailKey: "customer-email:",
	}
}

func (c *customerCache) ReadCacheByID(ctx context.Context, customerID string) (*entity.Customer, error) {
	c.logger.Debug("Getting customer cache", "customerID", customerID, "traceID", ctx.Value("traceID"))
	key := c.idKey + customerID

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

func (c *customerCache) WriteCacheByID(ctx context.Context, customer entity.Customer) error {
	c.logger.Debug("Creating customer cache", "data", customer, "traceID", ctx.Value("traceID"))
	key := c.idKey + *customer.ID

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

func (c *customerCache) ReadCacheByEmail(ctx context.Context, customerEmail string) (*entity.Customer, error) {
	c.logger.Debug("Getting customer cache by email", "customerEmail", customerEmail, "traceID", ctx.Value("traceID"))
	key := c.emailKey + customerEmail

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

func (c *customerCache) WriteCacheByEmail(ctx context.Context, customer entity.Customer) error {
	c.logger.Debug("Creating customer cache by email", "data", customer, "traceID", ctx.Value("traceID"))
	key := c.emailKey + customer.Email

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
