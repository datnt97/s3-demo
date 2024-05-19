package handler

import (
	"net/http"
	"tronglv_upload_svc/helper/server/http/response"
	"tronglv_upload_svc/internal/registry"
)

type UploadHandler interface {
	UploadFileS3() http.HandlerFunc
}

type uploadHandler struct {
	reg *registry.ServiceContext
}

func NewUploadHandler(reg *registry.ServiceContext) UploadHandler {
	return &uploadHandler{
		reg: reg,
	}
}

func (p *uploadHandler) UploadFileS3() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		// response.Error(r.Context(), w, err)
		// var data []string
		response.OkJson(r.Context(), w, "Upload Success", nil)

	}
}
