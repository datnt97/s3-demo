package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"tronglv_upload_svc/helper/errors"
	"tronglv_upload_svc/helper/s3"
	"tronglv_upload_svc/helper/server/http/response"
	"tronglv_upload_svc/helper/util"
	"tronglv_upload_svc/internal/registry"
	"tronglv_upload_svc/internal/service"
	"tronglv_upload_svc/internal/types/request"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type UploadHandler interface {
	UploadFileS3() http.HandlerFunc
	UploadFileS3Presign() http.HandlerFunc
}

type uploadHandler struct {
	reg       *registry.ServiceContext
	uploadSvc service.UploadService
}

func NewUploadHandler(reg *registry.ServiceContext) UploadHandler {
	return &uploadHandler{
		reg:       reg,
		uploadSvc: service.NewUploadService(reg),
	}
}
func (p *uploadHandler) parseFileUpload(r *http.Request) ([]*request.FileInfo, error) {
	fu := s3.NewFileUpload(r)
	files, err := fu.Parse("images[]")
	if err != nil {

		return nil, err
	}

	var items []*request.FileInfo
	for _, val := range files {
		f, err := val.Open()
		if err != nil {
			logx.Error(err)
			continue
		}

		fb, e := io.ReadAll(f)
		if e != nil {
			_ = f.Close()
			continue
		}
		_ = f.Close()

		items = append(items, &request.FileInfo{
			FileName: fmt.Sprintf("%s-%s", stringx.Randn(12), val.Filename),
			FileData: fb,
			FileSize: val.Size,
		})
	}

	if len(items) == 0 {

		return nil, errors.BadRequest(fmt.Errorf("missing images"))
	}
	return items, nil
}

func (p *uploadHandler) UploadFileS3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		items, err := p.parseFileUpload(r)
		if err != nil {
			response.Error(r.Context(), w, errors.BadRequest(err))
			return
		}

		resp, err := p.uploadSvc.UploadS3(r.Context(), &request.UploadAttachmentRequest{
			ServiceName: r.FormValue("service_name"),
			Acl:         util.String("public-read"),
			Attachments: items,
		})
		if err != nil {

			response.Error(r.Context(), w, err)
			return
		}

		response.OkJson(r.Context(), w, resp, nil)

	}
}

func (p *uploadHandler) UploadFileS3Presign() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := p.parseFileUpload(r)
		if err != nil {
			response.Error(r.Context(), w, errors.BadRequest(err))
			return
		}

		resp, err := p.uploadSvc.UploadS3(r.Context(), &request.UploadAttachmentRequest{
			ServiceName: r.FormValue("service_name"),
			Acl:         util.String("public-read"),
			Attachments: items,
			IsPresigned: util.Bool(true),
			Duration:    util.Duration(1 * time.Minute),
		})
		if err != nil {

			response.Error(r.Context(), w, err)
			return
		}

		response.OkJson(r.Context(), w, resp, nil)

	}
}
