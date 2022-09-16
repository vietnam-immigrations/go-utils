package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"
	vscontext "github.com/vietnam-immigrations/go-utils/v2/pkg/context"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
)

const (
	stageAll = "all"
)

// GetCommonParameter returns common parameter (stage = "all")
func GetCommonParameter(ctx context.Context, namespace, name string, decryption bool) (string, error) {
	return getParameter(ctx, stageAll, namespace, name, decryption)
}

// GetStageParameter returns ssm parameter. Stage must be in context.
func GetStageParameter(ctx context.Context, namespace, name string, decryption bool) (string, error) {
	stage, ok := ctx.Value(vscontext.KeyStage).(string)
	if !ok {
		return "", fmt.Errorf("missing stage in context")
	}

	return getParameter(ctx, stage, namespace, name, decryption)
}

func getParameter(ctx context.Context, stage, namespace, name string, decryption bool) (string, error) {
	log := logger.FromContext(ctx)

	ssmKey := fmt.Sprintf("/%s/%s%s", namespace, stage, name)
	log.Infof("get [%s] variable", ssmKey)
	ssmClient, err := newClient(ctx)
	if err != nil {
		log.Errorf("failed to create SSM client: %s", err)
		return "", err
	}

	getParameter, err := ssmClient.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(ssmKey),
		WithDecryption: aws.Bool(decryption),
	})
	if err != nil {
		log.Errorf("failed to read SSM: %s", err)
		return "", errors.Wrap(err, fmt.Sprintf("could not find ssm parameter: %s", ssmKey))
	}
	log.Infof("found in SSM")
	return *getParameter.Parameter.Value, nil
}
