package mongodb

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/vietnam-immigrations/go-utils/pkg/aws/ssm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	initClient sync.Once
	client     *mongo.Client
	db         string
)

func newClient(ctx context.Context, log *logrus.Entry, stage string) (*mongo.Client, error) {
	initClient.Do(func() {
		log.Infof("init mongodb client")
		mongoHost, err := ssm.GetParameter(ctx, log, "vs2", "all", "/mongo/host", false)
		if err != nil {
			return
		}
		mongoUsername, err := ssm.GetParameter(ctx, log, "vs2", stage, "/mongo/username", false)
		if err != nil {
			return
		}
		mongoPassword, err := ssm.GetParameter(ctx, log, "vs2", stage, "/mongo/password", true)
		if err != nil {
			return
		}
		mongoDatabase, err := ssm.GetParameter(ctx, log, "vs2", stage, "/mongo/db", false)
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

func Disconnect(ctx context.Context, log *logrus.Entry) {
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
