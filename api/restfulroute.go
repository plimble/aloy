package api

import (
	"github.com/plimble/aloy/aloy"
	"github.com/plimble/aloy/api/mw"
	"github.com/plimble/aloy/api/restful"
	"gopkg.in/kataras/iris.v6"
)

func restfulroute(srv *iris.Framework, uc aloy.UsecaseInterface) {
	srv.Post("/webhook/github", mw.H(restful.WebhookGithub, uc))
	srv.Post("/webhook/gitlab", mw.H(restful.WebhookGitlab, uc))
}
