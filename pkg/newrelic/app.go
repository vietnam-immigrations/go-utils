package newrelic

import (
	"context"
	"sync"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/aws/ssm"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
)

var (
	initApp sync.Once
	app     *newrelic.Application
)

func NewApplication(ctx context.Context) (*newrelic.Application, error) {
	log := logger.FromContext(ctx)
	var err error
	initApp.Do(func() {
		log.Infof("init newrelic app")

		name, ssmErr := ssm.GetStageParameter(ctx, "vs2", "/newrelic/name", false)
		if err != nil {
			err = ssmErr
			return
		}
		log.Infof("newrelic app [%s]", name)
		license, ssmErr := ssm.GetStageParameter(ctx, "vs2", "/newrelic/license", true)
		if err != nil {
			err = ssmErr
			return
		}

		app, err = newrelic.NewApplication(
			newrelic.ConfigAppName(name),
			newrelic.ConfigLicense(license),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)

		if err == nil {
			log.Infof("try to connect to newrelic")
			err = app.WaitForConnection(10 * time.Second)
		}
	})
	if err != nil {
		log.Errorf("failed to create newrelic app: %s", err)
		return nil, err
	}
	log.Infof("newrelic app created")
	return app, nil
}

func Shutdown(ctx context.Context) {
	log := logger.FromContext(ctx)
	if app != nil {
		log.Infof("shutdown newrelic")
		app.Shutdown(10 * time.Second)
	}
}
