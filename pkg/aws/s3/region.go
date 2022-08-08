package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

func NewClientForRegion(ctx context.Context, log *logrus.Entry, region string) (*s3.Client, error) {
	log.Infof("create s3 client for region [%s]", region)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Errorf("failed to load default config: %s", err)
		return nil, err
	}
	cfg.Region = region
	return s3.NewFromConfig(cfg), nil
}
