package mongodb

import (
	"context"

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
