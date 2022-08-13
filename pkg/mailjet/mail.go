package mailjet

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/pkg/errors"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/aws/ssm"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/retry"
)

// Send sends email
func Send(ctx context.Context, m mailjet.InfoMessagesV31) error {
	log := logger.FromContext(ctx)

	username, err := ssm.GetCommonParameter(ctx, "vs2", "/mailjet/username", false)
	if err != nil {
		return errors.Wrap(err, "failed to get mailjet username")
	}
	password, err := ssm.GetCommonParameter(ctx, "vs2", "/mailjet/password", true)
	if err != nil {
		return errors.Wrap(err, "failed to get mailjet password")
	}

	// TODO: mailjet should provide a way to use custom http.Client
	// see https://github.com/mailjet/mailjet-apiv3-go/issues/95
	http.DefaultClient.Timeout = 60 * time.Second
	client := mailjet.NewMailjetClient(username, password)
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{m},
	}

	return retry.Do(ctx, func() error {
		res, err := client.SendMailV31(&messages)
		if err != nil {
			log.Errorf("failed to send email [%s]: %s", m.Subject, err)
			return errors.Wrap(err, "failed to send email")
		}
		log.Infof("[%d]sent status: %s", len(res.ResultsV31), res.ResultsV31[0].Status)
		if res.ResultsV31[0].Status != "success" {
			log.Errorf("failed to send email [%s] to [%+v]: %s", m.Subject, m.To, res.ResultsV31[0].Status)
			return fmt.Errorf("failed to send email [%s] to [%+v]: %s", m.Subject, m.To, res.ResultsV31[0].Status)
		}
		log.Infof("email [%s] sent to [%+v]", m.Subject, m.To)
		return nil
	}, 3, 100*time.Millisecond)
}
