package cv

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awssns "github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/dchest/uniuri"
	"github.com/sirupsen/logrus"
	"github.com/vietnam-immigrations/go-utils/pkg/aws/sns"
	"github.com/vietnam-immigrations/go-utils/pkg/aws/ssm"
	"github.com/vietnam-immigrations/go-utils/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadToS3AndSendSNS(ctx context.Context, log *logrus.Entry, stage string, fileNames []string, fileContents []io.ReadCloser) error {
	s3Bucket, err := ssm.GetParameter(ctx, log, "vs2", stage, "/result/s3BucketName", false)
	if err != nil {
		log.Errorf("failed to get s3 location: %s", err)
		return err
	}

	result := &mongodb.Result{
		ID:        primitive.NewObjectID(),
		S3DirKey:  uniuri.NewLen(12),
		Files:     make([]mongodb.ResultFile, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// copy files to S3
	for i, fileName := range fileNames {
		_, err := putToS3Bucket(ctx, log, result, s3Bucket, fileName, fileContents[i])
		if err != nil {
			log.Errorf("failed to put to S3: %s", err)
			return err
		}
		result.Files = append(result.Files, mongodb.ResultFile{
			Name:         fileName,
			Processed:    false,
			ErrorMessage: "",
			OrderNumber:  "",
		})
	}

	// save result to mongo
	colResult, err := mongodb.CollectionResult(ctx, log, stage)
	if err != nil {
		log.Errorf("failed to get mongodb collection: %s", err)
		return err
	}
	_, err = colResult.InsertOne(ctx, result)
	if err != nil {
		log.Errorf("failed to insert result to mongo: %s", err)
		return err
	}

	// publish SNS
	snsClient, err := sns.NewClient(ctx, log)
	if err != nil {
		log.Errorf("failed to create SNS client: %s", err)
		return err
	}
	newResultTopic, err := ssm.GetParameter(ctx, log, "vs2", stage, "/sns/newResult/arn", false)
	if err != nil {
		log.Errorf("failed to get SNS topic arn: %s", err)
	}
	log.Infof("messages will be published to [%s]", newResultTopic)

	for i, fileName := range fileNames {
		_, err := snsClient.Publish(ctx, &awssns.PublishInput{
			Message: aws.String(fmt.Sprintf("New result file [%s]", fileName)),
			MessageAttributes: map[string]types.MessageAttributeValue{
				"filename": {
					DataType:    aws.String("String"),
					StringValue: aws.String(fileName),
				},
				"resultId": {
					DataType:    aws.String("String"),
					StringValue: aws.String(result.ID.Hex()),
				},
			},
			MessageDeduplicationId: aws.String(fmt.Sprintf("%s-%d", result.ID.Hex(), i)),
			MessageGroupId:         aws.String(stage),
			Subject:                aws.String("New result file"),
			TopicArn:               aws.String(newResultTopic),
		})
		if err != nil {
			log.Warnf("failed to send SNS result file uploaded for [%s]: %s", fileName, err)
			// TODO: set processing to failed for this file?
		}
	}

	return nil
}
