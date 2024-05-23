package service

import (
	"bytes"
	"context"
	"fmt"
	"tronglv_upload_svc/helper/s3"
	"tronglv_upload_svc/helper/util"

	"tronglv_upload_svc/internal/registry"
	"tronglv_upload_svc/internal/types/request"
	"tronglv_upload_svc/internal/types/response"
)

type UploadService interface {
	UploadS3(ctx context.Context, req *request.UploadAttachmentRequest) ([]*response.UploadResp, error)
}

type uploadSvcImpl struct {
	reg           *registry.ServiceContext
	uploadStorage s3.S3Storage
}

func NewUploadService(reg *registry.ServiceContext) UploadService {
	return &uploadSvcImpl{
		reg:           reg,
		uploadStorage: reg.UploadStorage,
	}
}

func (s *uploadSvcImpl) UploadS3(ctx context.Context, req *request.UploadAttachmentRequest) ([]*response.UploadResp, error) {

	var items []*response.UploadResp

	for _, file := range req.Attachments {
		fileName := fmt.Sprintf("%s/%s", util.Slug(req.ServiceName), file.FileName)
		resp, err := s.uploadStorage.Upload(ctx, fileName, bytes.NewReader(file.FileData), req.Acl)
		if err != nil {
			return nil, err
		}

		fmt.Println("vao trong nay 2 resp ", resp)
		items = append(items, &response.UploadResp{
			Url: resp.FileUrl,
		})
	}

	return items, nil
}
