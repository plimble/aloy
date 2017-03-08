package aloy

import (
	"github.com/plimble/aloy/services/testrunner"
)

type GetGithubWebhookInput struct {
	Payload []byte
}

func (uc *Usecase) GetGithubWebhook(input *GetGithubWebhookInput) error {
	wh, err := uc.webhook.ParseGithubWebhook(input.Payload)
	uc.testrunner.Enqueue(testrunner.Message(*wh))

	return err
}
