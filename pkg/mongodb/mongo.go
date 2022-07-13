package mongodb

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func collection(ctx context.Context, log *logrus.Entry, stage string, name string) (*mongo.Collection, error) {
	client, err := newClient(ctx, log, stage)
	if err != nil {
		return nil, err
	}
	return client.Database(db).Collection(name), nil
}
