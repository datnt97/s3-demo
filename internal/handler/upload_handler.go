package handler

import (
	"fmt"
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
		ctx := r.Context()
		// response.Error(r.Context(), w, err)
		// var data []string
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			response.Error(ctx, w, err)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			response.Error(ctx, w, err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		response.OkJson(r.Context(), w, "Upload Success", nil)

	}
}
