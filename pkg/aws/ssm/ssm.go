package ssm

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetParameter(ctx context.Context, log *logrus.Entry, namespace, stage, name string, decryption bool) (string, error) {
	ssmKey := fmt.Sprintf("/%s/%s%s", namespace, stage, name)
	log.Infof("get [%s] variable", ssmKey)
	if fromEnvVar := getFromEnvVar(log, namespace, stage, name); fromEnvVar != "" {
		log.Infof("found from env variable")
		return fromEnvVar, nil
	}
	client, err := newClient(ctx, log)
	if err != nil {
		log.Errorf("failed to create SSM client: %s", err)
		return "", err
	}

	getParameter, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(ssmKey),
		WithDecryption: decryption,
	})
	if err != nil {
		log.Errorf("failed to read SSM: %s", err)
		return "", errors.Wrap(err, fmt.Sprintf("could not find ssm parameter: %s", ssmKey))
	}
	log.Infof("found in SSM")
	return *getParameter.Parameter.Value, nil
}

func getFromEnvVar(log *logrus.Entry, namespace, stage, name string) string {
	name = strings.ReplaceAll(name, "/", "_")
	envKey := fmt.Sprintf("SSM_%s_%s%s", strings.ToUpper(namespace), strings.ToUpper(stage), strings.ToUpper(name))
	log.Infof("looking for env var [%s]", envKey)
	return os.Getenv(envKey)
}
