package handler

import (
	"net/http"
	"tronglv_upload_svc/helper/server/http/handler"
	"tronglv_upload_svc/internal/registry"

	"github.com/zeromicro/go-zero/rest"
)

const (
	BasePrefix = "/upload-svc"
	RestPrefix = BasePrefix + "/api/v1"
)

type RestHandler struct {
	svc *registry.ServiceContext
}

func NewRestHandler(svc *registry.ServiceContext) RestHandler {
	return RestHandler{svc: svc}
}

func (h RestHandler) Register(svr *rest.Server) {
	handler.RegisterSwaggerHandler(svr, BasePrefix)
	registerUploadHandler(svr, h.svc)
}
func registerUploadHandler(svr *rest.Server, svc *registry.ServiceContext) {
	h := NewUploadHandler(svc)
	var path = "/uploadS3"
	svr.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    path,
					Handler: h.UploadFileS3(),
				},
			}...,
		),
		rest.WithPrefix(RestPrefix),
	)
}
