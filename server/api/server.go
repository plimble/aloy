package api

// import (
// 	"github.com/plimble/aloy/api/services"
// 	"github.com/plimble/aloy/services/config"
// 	"gopkg.in/kataras/iris.v6"
// 	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
// 	"gopkg.in/kataras/iris.v6/middleware/recover"
// )

// func New(cfg *config.Config) *iris.Framework {
// 	s, err := services.New(cfg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	srv := iris.New()

// 	srv.Adapt(httprouter.New())

// 	srv.Use(recover.New())

// 	restfulroute(srv, s)

// 	return srv
// }
