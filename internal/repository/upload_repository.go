package repository

import "context"

type UploadRepository interface {
	UploadS3(ctx context.Context) error
}

type uploadRepo struct {
}

func NewUploadRepository() UploadRepository {
	return &uploadRepo{}
}

func (p *uploadRepo) UploadS3(ctx context.Context) error {
	return nil
}
