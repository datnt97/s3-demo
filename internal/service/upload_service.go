package service

import (
	"bytes"
	"context"
	"fmt"
	"time"
	"tronglv_upload_svc/helper/s3"
	"tronglv_upload_svc/helper/util"

	"tronglv_upload_svc/internal/registry"
	"tronglv_upload_svc/internal/types/request"
	"tronglv_upload_svc/internal/types/response"
)

type UploadService interface {
	UploadS3(ctx context.Context, req *request.UploadAttachmentRequest) ([]*response.UploadResp, error)
	UploadS3CloudFront(ctx context.Context, req *request.UploadAttachmentRequest) ([]*response.UploadResp, error)
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
	var keys []string
	var items []*response.UploadResp

	for _, file := range req.Attachments {
		fileName := fmt.Sprintf("%s/%s", util.Slug(req.ServiceName), file.FileName)
		resp, err := s.uploadStorage.Upload(ctx, fileName, bytes.NewReader(file.FileData), false)
		if err != nil {
			return nil, err
		}

		keys = append(keys, resp.FileName)
		items = append(items, &response.UploadResp{
			Url: resp.FileUrl,
		})
	}

	if req.IsPresigned != nil && *req.IsPresigned {

		var expiration = 15 * time.Minute
		if req.Duration != nil {
			expiration = *req.Duration
		}
		urls, err := s.withPreSignedUrls(ctx, keys, expiration)
		if err != nil {
			return nil, err
		}

		// clear items
		items = nil

		for _, val := range urls {
			items = append(items, &response.UploadResp{
				Url: val,
			})
		}

	}

	return items, nil
}

func (s *uploadSvcImpl) UploadS3CloudFront(ctx context.Context, req *request.UploadAttachmentRequest) ([]*response.UploadResp, error) {
	var keys []string
	var items []*response.UploadResp

	for _, file := range req.Attachments {
		fileName := fmt.Sprintf("%s/%s", util.Slug(req.ServiceName), file.FileName)
		resp, err := s.uploadStorage.Upload(ctx, fileName, bytes.NewReader(file.FileData), true)
		if err != nil {
			return nil, err
		}

		keys = append(keys, resp.FileUrl)
		items = append(items, &response.UploadResp{
			Url: resp.FileUrl,
		})
	}

	if req.IsPresigned != nil && *req.IsPresigned {

		var expiration = 15 * time.Minute
		if req.Duration != nil {
			expiration = *req.Duration
		}
		urls, err := s.withCloudFrontSignedUrls(ctx, keys, expiration)
		if err != nil {
			return nil, err
		}

		// clear items
		items = nil

		for _, val := range urls {
			items = append(items, &response.UploadResp{
				Url: val,
			})
		}

	}

	return items, nil
}

func (s *uploadSvcImpl) withPreSignedUrls(ctx context.Context, objectKeys []string, expiration time.Duration) ([]string, error) {
	var urls []string
	for _, v := range objectKeys {

		url, e := s.uploadStorage.SignedUrl(ctx, v, expiration)
		if e != nil {
			return nil, e
		}
		urls = append(urls, *url)
	}
	return urls, nil
}

func (s *uploadSvcImpl) withCloudFrontSignedUrls(ctx context.Context, objectKeys []string, expiration time.Duration) ([]string, error) {
	var urls []string
	for _, v := range objectKeys {

		url, e := s.uploadStorage.CloundFrontSignUrl(ctx, v, expiration)
		if e != nil {
			return nil, e
		}
		urls = append(urls, *url)
	}
	return urls, nil
}
