package aloy

import (
	"github.com/plimble/aloy/services/testrunner"
)

type GetGitlabWebhookInput struct {
	Payload []byte
}

func (uc *Usecase) GetGitlabWebhook(input *GetGitlabWebhookInput) error {
	wh, err := uc.webhook.ParseGitlabWebhook(input.Payload)
	uc.testrunner.Enqueue(testrunner.Message(*wh))

	return err
}
