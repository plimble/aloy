package aloy

import (
	"github.com/plimble/aloy/services/badge"
	"github.com/plimble/aloy/services/config"
	"github.com/plimble/aloy/services/testrunner"
	"github.com/plimble/aloy/services/webhook"
)

type Services struct {
	config     *config.Config
	webhook    webhook.ServiceInterface
	testrunner testrunner.ServiceInterface
	badge      badge.ServiceInterface
}

func (s *Services) SetWebhook(service webhook.ServiceInterface) {
	s.webhook = service
}

func (s *Services) SetConfig(service *config.Config) {
	s.config = service
}

func (s *Services) SetTestRunner(service testrunner.ServiceInterface) {
	s.testrunner = service
}

func (s *Services) SetBadge(service badge.ServiceInterface) {
	s.badge = service
}
