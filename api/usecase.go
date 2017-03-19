package api

import (
	"github.com/plimble/aloy/aloy"
	"github.com/plimble/aloy/badge"
	"github.com/plimble/aloy/config"
	"github.com/plimble/aloy/testrunner"
	"github.com/plimble/aloy/webhook"
)

func NewUsecase(cfg *config.Config) aloy.UsecaseInterface {
	s := &aloy.Services{}

	s.SetConfig(cfg)
	s.SetWebhook(webhook.NewService())
	s.SetTestRunner(testrunner.NewService(testrunner.Options{
		MaxQueue:   cfg.MaxQueue,
		MaxRunner:  cfg.MaxRunner,
		GoTestTags: cfg.GoTestTags,
	}))
	s.SetBadge(badge.NewService())

	return aloy.New(s)
}
