package notification

import (
	"context"

	"github.com/pusher/pusher-http-go/v5"
	"github.com/sirupsen/logrus"
	"github.com/vietnam-immigrations/go-utils/pkg/mongodb"
)

const (
	channelGlobal     = "global"
	eventNotification = "notification"
)

type Notification struct {
	ID         string `json:"id"`
	CSSClasses string `json:"cssClasses"`
	Message    string `json:"message"`
}

func Create(ctx context.Context, log *logrus.Entry, stage string, item Notification) error {
	log.Infof("publish message: %+v", item)
	globalConfig, err := mongodb.GetConfig(ctx, log, stage)
	if err != nil {
		log.Errorf("failed to load global config: %s", err)
		return err
	}
	client := pusher.Client{
		AppID:   globalConfig.PusherAppID,
		Key:     globalConfig.PusherAppKey,
		Secret:  globalConfig.PusherAppSecret,
		Cluster: globalConfig.PusherCluster,
	}

	err = client.Trigger(channelGlobal, eventNotification, item)
	if err != nil {
		log.Errorf("failed to send message to pusher: %s", err)
		return err
	}
	log.Infof("message sent")
	return nil
}
