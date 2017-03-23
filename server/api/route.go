package api

import (
	"github.com/plimble/aloy/aloy"
	"github.com/plimble/aloy/server/api/handler"
	"gopkg.in/kataras/iris.v6"
)

func Route(srv *iris.Framework, a aloy.Aloy) {
	h := handler.New(a)

	srv.Post("/webhook", W(h.WebhookCallback))
}
