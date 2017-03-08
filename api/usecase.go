package api

import (
	"github.com/plimble/aloy/services/aloy"
	"github.com/plimble/aloy/services/config"
	"github.com/plimble/aloy/services/queue"
	"github.com/plimble/aloy/services/webhook"
)

func NewUsecase() aloy.UsecaseInterface {
	s := &aloy.Services{}

	s.SetConfig(config.Get())
	s.SetWebhook(webhook.New())
	s.SetQueue(queue.New())

	return aloy.New(s)
}
