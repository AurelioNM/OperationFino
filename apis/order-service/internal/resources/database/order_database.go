package database

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"cmd/order-service/internal/metrics"
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderGateway struct {
	logger  slog.Logger
	metrics *metrics.OrderMetrics
	db      *mongo.Client
}

func NewOrderGateway(l slog.Logger, m *metrics.OrderMetrics, db *mongo.Client) gateway.OrderGateway {
	return &orderGateway{
		logger:  *l.With("layer", "order-database"),
		metrics: m,
		db:      db,
	}
}

func (g *orderGateway) GetOrderByID(ctx context.Context, orderID *string) (*entity.Order, error) {
	g.logger.Debug("Getting order by ID from DB", "ID", orderID, "traceID", ctx.Value("traceID"))
	collection := g.db.Database("order-service").Collection("order")

	var order entity.Order
	filter := bson.M{"id": orderID}
	start := time.Now()

	err := collection.FindOne(ctx, filter).Decode(&order)
	g.metrics.MeasureExternalDuration(start, "database", "OrderDB", "GetOrderByID", "")
	if err != nil {
		if err != mongo.ErrNoDocuments {
			g.logger.Error("Order not found by ID", "error", err, "traceID", ctx.Value("traceID"))
			return nil, err
		}
		g.logger.Error("Failed to find order by ID in DB", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return &order, nil
}

func (g *orderGateway) GetOrdersByCustomerID(ctx context.Context, customerID *string) ([]*entity.Order, error) {
	g.logger.Debug("Getting orders list by customerID from DB", "customerID", customerID, "traceID", ctx.Value("traceID"))
	collection := g.db.Database("order-service").Collection("order")

	var orders []*entity.Order
	filter := bson.M{"customer.id": customerID}
	start := time.Now()

	cursor, err := collection.Find(ctx, filter)
	g.metrics.MeasureExternalDuration(start, "database", "OrderDB", "GetOrdersByCustomerID", "")
	if err != nil {
		g.logger.Error("Failed to find orders list by customerID in DB", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &orders); err != nil {
		g.logger.Error("Failed decode orders list", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	return orders, nil
}

func (g *orderGateway) CreateOrder(ctx context.Context, order *entity.Order) (*string, error) {
	g.logger.Debug("Inserting order into DB", "traceID", ctx.Value("traceID"))
	collection := g.db.Database("order-service").Collection("order")
	start := time.Now()

	insertResult, err := collection.InsertOne(ctx, order)
	g.metrics.MeasureExternalDuration(start, "database", "OrderDB", "CreateOrder", "")
	if err != nil {
		g.logger.Error("Failed to insert order into DB", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	g.logger.Info("Sucessfull insert of order into DB", "insertedID", insertResult.InsertedID, "traceID", ctx.Value("traceID"))
	return order.ID, nil
}

func (g *orderGateway) DeleteOrderByID(ctx context.Context, orderID *string) error {
	g.logger.Debug("Deleting order by ID from DB", "ID", orderID, "traceID", ctx.Value("traceID"))
	collection := g.db.Database("order-service").Collection("order")

	filter := bson.M{"id": orderID}
	start := time.Now()

	deletedResult, err := collection.DeleteOne(ctx, filter)
	g.metrics.MeasureExternalDuration(start, "database", "OrderDB", "DeleteOrderByID", "")
	if err != nil {
		g.logger.Error("Failed to find order by ID in DB", "error", err, "traceID", ctx.Value("traceID"))
		return err
	}

	if deletedResult.DeletedCount == 0 {
		g.logger.Error("Order not found by ID", "error", err, "traceID", ctx.Value("traceID"))
		return fmt.Errorf("Order not found with ID=%s", *orderID)
	}

	return nil
}
