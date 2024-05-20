package service

import (
	"context"
	"tronglv_upload_svc/internal/registry"
	"tronglv_upload_svc/internal/repository"
)

type UploadService interface {
	UploadS3(ctx context.Context) error
}

type uploadSvcImpl struct {
	reg        *registry.ServiceContext
	uploadRepo repository.UploadRepository
}

func NewUploadService(reg *registry.ServiceContext) UploadService {
	return &uploadSvcImpl{
		reg:        reg,
		uploadRepo: reg.UploadRepo,
	}
}

func (s *uploadSvcImpl) UploadS3(ctx context.Context) error {
	return nil
}
