package mailjet

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/aws/ssm"
)

func Send(ctx context.Context, log *logrus.Entry, m mailjet.InfoMessagesV31) error {
	username, err := ssm.GetParameter(ctx, log, "vs2", "all", "/mailjet/username", false)
	if err != nil {
		return errors.Wrap(err, "failed to get mailjet username")
	}
	password, err := ssm.GetParameter(ctx, log, "vs2", "all", "/mailjet/password", true)
	if err != nil {
		return errors.Wrap(err, "failed to get mailjet password")
	}

	http.DefaultClient.Timeout = 30 * time.Second
	client := mailjet.NewMailjetClient(username, password)
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{m},
	}
	res, err := client.SendMailV31(&messages)
	if err != nil {
		log.Errorf("failed to send email [%s]: %s", m.Subject, err)
		return errors.Wrap(err, "failed to send email")
	}
	log.Infof("[%d]sent status: %s", len(res.ResultsV31), res.ResultsV31[0].Status)
	if res.ResultsV31[0].Status != "success" {
		log.Errorf("failed to send email [%s] to [%v]: %s", m.Subject, m.To, res.ResultsV31[0].Status)
		return fmt.Errorf("failed to send email [%s] to [%v]: %s", m.Subject, m.To, res.ResultsV31[0].Status)
	}

	log.Infof("email [%s] sent to [%v]", m.Subject, m.To)
	return nil
}
