package main

import (
	"fmt"

	"github.com/plimble/aloy/aloy"
	"github.com/plimble/aloy/server/api"
	"github.com/plimble/goconfig"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/recover"
)

func main() {
	appConfig := &Config{}
	aloyConfig := &aloy.Config{}
	goconfig.MustProcess("aloy", appConfig)
	goconfig.MustProcess("aloy", aloyConfig)

	srv := iris.New()
	srv.Adapt(httprouter.New())
	srv.Use(recover.New())

	a, err := aloy.New(aloyConfig)
	if err != nil {
		panic(err)
	}

	api.Route(srv, a)

	fmt.Println("server start on", appConfig.Addr)
	srv.Listen(appConfig.Addr)
}
