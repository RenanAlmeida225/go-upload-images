package config

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	clientS3 *s3.Client
)

func Init() error {
	err := InitializeGodotenv()
	if err != nil {
		return err
	}
	clientS3, err = InitializeS3()
	if err != nil {
		return err
	}
	return nil
}

func GetClientS3() *s3.Client {
	return clientS3
}
