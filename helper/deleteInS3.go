package helper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteInS3(name string) error {
	_, err := clientS3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String("upload-images-go"),
		Key:    aws.String(name),
	})

	if err != nil {
		return err
	}

	return nil
}
