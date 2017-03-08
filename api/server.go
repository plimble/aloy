package api

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/recover"
)

func New() *iris.Framework {
	uc := NewUsecase()
	srv := iris.New()

	srv.Adapt(httprouter.New())

	srv.Use(recover.New())

	restfulroute(srv, uc)

	return srv
}
