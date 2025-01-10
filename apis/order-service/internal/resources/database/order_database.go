package database

import (
	"cmd/order-service/internal/domain/entity"
	"cmd/order-service/internal/domain/gateway"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderGateway struct {
	logger slog.Logger
	db     *mongo.Client
}

func NewOrderGateway(l slog.Logger, db *mongo.Client) gateway.OrderGateway {
	return &orderGateway{
		logger: *l.With("layer", "order-database"),
		db:     db,
	}
}

func (g *orderGateway) GetOrderByID(ctx context.Context, orderID *string) (*entity.Order, error) {
	g.logger.Debug("Getting order by ID from DB", "ID", orderID, "traceID", ctx.Value("traceID"))
	collection := g.db.Database("order-service").Collection("order")

	var order entity.Order
	filter := bson.M{"id": orderID}

	err := collection.FindOne(ctx, filter).Decode(&order)
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

	cursor, err := collection.Find(ctx, filter)
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

	insertResult, err := collection.InsertOne(ctx, order)
	if err != nil {
		g.logger.Error("Failed to insert order into DB", "error", err, "traceID", ctx.Value("traceID"))
		return nil, err
	}

	g.logger.Info("Sucessfull insert of order into DB", "insertedID", insertResult.InsertedID, "traceID", ctx.Value("traceID"))
	return order.ID, nil
}
