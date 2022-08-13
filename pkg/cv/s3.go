package cv

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/aws/s3"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/logger"
	"github.com/vietnam-immigrations/go-utils/v2/pkg/mongodb"
)

func putToS3Bucket(ctx context.Context, result *mongodb.Result, bucket string, fileName string, fileContent io.ReadCloser) (string, error) {
	log := logger.FromContext(ctx)
	log.Infof("upload file [%s] to [%s]", fileName, bucket)

	checkContent, err := ioutil.ReadAll(fileContent)
	if err != nil {
		log.Errorf("failed to read []byte from [%s]", fileName)
		return "", err
	}
	defer func(fileContent io.ReadCloser) {
		err := fileContent.Close()
		if err != nil {
			log.Warnf("failed to close file [%s]", fileName)
		}
	}(fileContent)
	newContent := bytes.NewReader(checkContent)

	s3Client, err := s3.NewClient(ctx)
	if err != nil {
		log.Errorf("failed to initiate s3 client: %s", err)
		return "", err
	}

	key := fmt.Sprintf("%s/%s", result.S3DirKey, fileName)
	_, err = s3Client.PutObject(ctx, &awss3.PutObjectInput{
		Body:   newContent,
		Key:    aws.String(key),
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Errorf("failed to put [%s] to S3 [%s] [%s]", fileName, bucket, key)
		return "", err
	}

	log.Infof("file [%s] upload to [%s] [%s]", fileName, bucket, key)
	return key, nil
}
