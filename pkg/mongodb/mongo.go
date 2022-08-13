package mongodb

import (
	"context"

	context2 "github.com/vietnam-immigrations/go-utils/v2/pkg/context"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// collection returns collection. Stage must be in context.
func collection(ctx context.Context, name string) (*mongo.Collection, error) {
	log := logger.FromContext(ctx)
	log.Infof("get collection [%s]", name)
	client, err := newClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database(db).Collection(name), nil
}

// AddOrderToContext adds order data to context
func AddOrderToContext(ctx context.Context, order Order) context.Context {
	result := context.WithValue(ctx, context2.KeyOrderID, order.ID)
	result = context.WithValue(result, context2.KeyOrderWooID, order.OrderID)
	result = context.WithValue(result, context2.KeyOrderNumber, order.Number)
	return result
}
