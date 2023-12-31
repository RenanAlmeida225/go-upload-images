package s3

import (
	"github.com/RenanAlmeida225/go-upload-images/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	clientS3 *s3.Client
)

func Init() {
	clientS3 = config.GetClientS3()
}
