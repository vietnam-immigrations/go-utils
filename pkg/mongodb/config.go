package mongodb

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionConfig(ctx context.Context, log *logrus.Entry, stage string) (*mongo.Collection, error) {
	return collection(ctx, log, stage, "config")
}

type Config struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	PusherAppID     string             `bson:"pusherAppId" json:"pusherAppId"`
	PusherAppKey    string             `bson:"pusherAppKey" json:"pusherAppKey"`
	PusherAppSecret string             `bson:"pusherAppSecret" json:"pusherAppSecret"`
	PusherCluster   string             `bson:"pusherCluster" json:"pusherCluster"`
}

func GetConfig(ctx context.Context, log *logrus.Entry, stage string) (*Config, error) {
	log.Infof("getting global configuration")
	colConfig, err := CollectionConfig(ctx, log, stage)
	if err != nil {
		return nil, err
	}
	findAll, err := colConfig.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	items := make([]Config, 0)
	err = findAll.All(ctx, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		log.Warnf("missing global configuration, creating a new one")
		globalConfig := Config{
			ID: primitive.NewObjectID(),
		}
		_, err := colConfig.InsertOne(ctx, globalConfig)
		if err != nil {
			return nil, err
		}
		return &globalConfig, nil
	}
	if len(items) > 1 {
		return nil, fmt.Errorf("too many global configurations objects [%d]", len(items))
	}
	globalConfig := items[0]
	return &globalConfig, nil
}