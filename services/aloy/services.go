package aloy

import (
	"github.com/plimble/aloy/services/config"
	"github.com/plimble/aloy/services/queue"
	"github.com/plimble/aloy/services/webhook"
)

type Services struct {
	config  *config.Config
	webhook webhook.ServiceInterface
	queue   queue.ServiceInterface
}

func (s *Services) SetWebhook(service webhook.ServiceInterface) {
	s.webhook = service
}

func (s *Services) SetConfig(service *config.Config) {
	s.config = service
}

func (s *Services) SetQueue(service queue.ServiceInterface) {
	s.queue = service
}
