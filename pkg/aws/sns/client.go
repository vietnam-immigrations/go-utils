package sns

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/sirupsen/logrus"
)

var (
	initClient sync.Once
	client     *sns.Client
)

func NewClient(ctx context.Context, log *logrus.Entry) (*sns.Client, error) {
	var err error
	initClient.Do(func() {
		cfg, e := config.LoadDefaultConfig(ctx)
		if e != nil {
			err = e
			return
		}
		client = sns.NewFromConfig(cfg)
	})
	if err != nil {
		return nil, err
	}
	log.Info("sns client created")
	return client, nil
}
