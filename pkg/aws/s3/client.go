package s3

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

var (
	initClient sync.Once
	client     *s3.Client
)

func NewClient(ctx context.Context, log *logrus.Entry) (*s3.Client, error) {
	var err error
	initClient.Do(func() {
		cfg, e := config.LoadDefaultConfig(ctx)
		if e != nil {
			err = e
			return
		}
		client = s3.NewFromConfig(cfg)
	})
	if err != nil {
		return nil, err
	}
	log.Info("s3 client created")
	return client, nil
}
