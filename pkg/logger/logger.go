package logger

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	vscontext "github.com/vietnam-immigrations/go-utils/v2/pkg/context"
)

type LogField string

var (
	initLogger sync.Once
	logger     *logrus.Logger
)

func getLogger() *logrus.Logger {
	initLogger.Do(func() {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
	})
	return logger
}

// FromContext returns logger for this context
func FromContext(ctx context.Context) *logrus.Entry {
	fields := make(logrus.Fields, 0)
	for _, logField := range vscontext.Keys {
		fields[logField] = ctx.Value(logField)
	}
	logger := getLogger()
	return logger.WithFields(fields)
}
