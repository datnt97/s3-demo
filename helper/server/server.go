package server

import (
	"tronglv_upload_svc/helper/server/http/middleware"

	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest"
)

func MustSetup(c Config) {
	Initialize()
	if !c.StatLog {
		stat.DisableLog()
	}
	if !c.LoadLog {
		load.DisableLog()
	}
}

func Initialize() {
	for _, v := range Providers() {
		v.Register()
	}
}

func NewHttpServer(c Config, h RestHandler, opts ...rest.RunOption) *rest.Server {
	MustSetup(c)
	srv := rest.MustNewServer(c.Http, opts...)
	srv.Use(middleware.NewRecoveryMiddleware(c.Env).Handle)
	h.Register(srv)
	return srv
}
