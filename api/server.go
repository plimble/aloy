package api

import (
	"github.com/plimble/aloy/config"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/recover"
)

func Run() {
	cfg := config.Get()
	uc := NewUsecase(cfg)
	srv := iris.New()

	srv.Adapt(httprouter.New())

	srv.Use(recover.New())

	restfulroute(srv, uc)

	srv.Listen(cfg.Addr)
}
