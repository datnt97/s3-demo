package registry

import "tronglv_upload_svc/internal/config"

type ServiceContext struct {
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{}
}
