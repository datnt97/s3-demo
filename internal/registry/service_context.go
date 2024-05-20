package registry

import (
	"tronglv_upload_svc/internal/config"
	"tronglv_upload_svc/internal/repository"
)

type ServiceContext struct {
	Config     config.Config
	UploadRepo repository.UploadRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UploadRepo: repository.NewUploadRepository(),
	}
}
