package mongodb

import (
	"context"
	"fmt"
	"sync"

	"github.com/vietnam-immigrations/go-utils/v2/pkg/aws/ssm"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	initClient sync.Once
	client     *mongo.Client
	db         string
)

// newClient returns mongo client. Stage must be in context.
func newClient(ctx context.Context) (*mongo.Client, error) {
	log := logger.FromContext(ctx)

	initClient.Do(func() {
		log.Infof("init mongodb client")
		mongoHost, err := ssm.GetStageParameter(ctx, "vs2", "/mongo/host", false)
		if err != nil {
			return
		}
		mongoUsername, err := ssm.GetStageParameter(ctx, "vs2", "/mongo/username", false)
		if err != nil {
			return
		}
		mongoPassword, err := ssm.GetStageParameter(ctx, "vs2", "/mongo/password", true)
		if err != nil {
			return
		}
		mongoDatabase, err := ssm.GetStageParameter(ctx, "vs2", "/mongo/db", false)
		if err != nil {
			return
		}
		log.Infof("host = %s, user = %s, db = %s", mongoHost, mongoUsername, mongoDatabase)
		mongoFullUrl := fmt.Sprintf("mongodb+srv://%s", mongoHost)

		c, err := mongo.NewClient(
			options.Client().ApplyURI(mongoFullUrl).SetAuth(options.Credential{
				Username: mongoUsername,
				Password: mongoPassword,
			}),
		)
		if err != nil {
			log.Errorf("failed to create mongodb client: %s", err)
			return
		}
		err = c.Connect(ctx)
		if err != nil {
			log.Errorf("failed to connect to mongodb: %s", err)
			return
		}

		client = c
		db = mongoDatabase
	})

	if client == nil {
		return nil, fmt.Errorf("failed to init mongodb client")
	}
	return client, nil
}

// Disconnect disconnects from db
func Disconnect(ctx context.Context) {
	log := logger.FromContext(ctx)
	if client != nil {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Errorf("failed to disconnect mongodb client: %s", err)
		} else {
			log.Infof("mongodb client disconnected")
			client = nil
			initClient = sync.Once{}
		}
	}
}
