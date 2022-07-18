package mongodb

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionResult(ctx context.Context, log *logrus.Entry, stage string) (*mongo.Collection, error) {
	return collection(ctx, log, stage, "results")
}

type ResultFile struct {
	Name        string `bson:"name" json:"name"`
	Processed   bool   `bson:"processed" json:"processed"`
	OrderNumber string `bson:"orderNumber" json:"orderNumber"`
}

type Result struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	S3DirKey string             `bson:"s3DirKey" json:"s3DirKey"`
	Files    []ResultFile       `bson:"files" json:"files"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
