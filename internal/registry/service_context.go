package registry

import (
	"tronglv_upload_svc/helper/s3"
	"tronglv_upload_svc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UploadStorage s3.S3Storage
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UploadStorage: s3.NewS3Storage(c.AwsS3),
	}
}
