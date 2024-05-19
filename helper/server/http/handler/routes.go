package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterSwaggerHandler(svr *rest.Server, prefix string) {
	h := NewSwaggerHandler()

	svr.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "swagger.yaml",
				Handler: h.File(prefix),
			},
			{
				Method:  http.MethodGet,
				Path:    "docs",
				Handler: h.Docs(prefix),
			},
		},
		rest.WithPrefix(prefix),
	)
}
