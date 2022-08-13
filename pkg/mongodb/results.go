package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionResult(ctx context.Context) (*mongo.Collection, error) {
	return collection(ctx, "results")
}

type ResultFile struct {
	Name         string `bson:"name" json:"name"`
	Processed    bool   `bson:"processed" json:"processed"`
	ErrorMessage string `bson:"errorMessage" json:"errorMessage"`
	OrderNumber  string `bson:"orderNumber" json:"orderNumber"`
	// PassportNumber used to match CV manually
	PassportNumber string `bson:"passportNumber" json:"passportNumber"`
}

type Result struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	S3DirKey string             `bson:"s3DirKey" json:"s3DirKey"`
	Files    []ResultFile       `bson:"files" json:"files"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
