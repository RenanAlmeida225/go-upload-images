package s3

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func SaveInS3(image *multipart.FileHeader, name string) (string, error) {
	f, err := image.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	uploader := manager.NewUploader(clientS3)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("upload-images-go"),
		Key:    aws.String(name),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {

		return "", err
	}

	return result.Location, nil
}
