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

var (
	initPreSignClient sync.Once
	preSignClient     *s3.PresignClient
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
		log.Info("s3 client created")
	})
	if err != nil {
		log.Errorf("failed to create s3 client: %s", err)
		return nil, err
	}
	return client, nil
}

func NewPreSignClient(ctx context.Context, log *logrus.Entry) (*s3.PresignClient, error) {
	var err error
	initPreSignClient.Do(func() {
		_, e := NewClient(ctx, log)
		if e != nil {
			err = e
			return
		}
		preSignClient = s3.NewPresignClient(client)
		log.Info("s3 pre sign client created")
	})
	if err != nil {
		log.Errorf("failed to create s3 pre sign client: %s", err)
		return nil, err
	}
	return preSignClient, nil
}
