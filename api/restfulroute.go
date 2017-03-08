package api

import (
	"github.com/plimble/aloy/api/restful"
	"github.com/plimble/aloy/services/aloy"
	"gopkg.in/kataras/iris.v6"
)

func restfulroute(srv *iris.Framework, uc aloy.UsecaseInterface) {
	srv.Post("/webhook/github", restful.WebhookGithub)
	srv.Post("/webhook/gitlab", restful.WebhookGitlab)
}
