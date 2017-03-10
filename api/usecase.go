package api

import (
	"github.com/plimble/aloy/services/aloy"
	"github.com/plimble/aloy/services/badge"
	"github.com/plimble/aloy/services/config"
	"github.com/plimble/aloy/services/testrunner"
	"github.com/plimble/aloy/services/webhook"
)

func NewUsecase(cfg *config.Config) aloy.UsecaseInterface {
	s := &aloy.Services{}

	s.SetConfig(cfg)
	s.SetWebhook(webhook.New())
	s.SetTestRunner(testrunner.New(testrunner.ServiceOptions{
		MaxQueue:       cfg.MaxQueue,
		MaxRunner:      cfg.MaxRunner,
		GithubUsername: cfg.GithubUsername,
		GithubPassword: cfg.GithubPassword,
		GitlabUsername: cfg.GitlabUsername,
		GitLabPassword: cfg.GitlabPassword,
		GoTestTags:     cfg.GoTestTags,
	}))
	s.SetBadge(badge.New())

	return aloy.New(s)
}
