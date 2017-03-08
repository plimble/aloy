package aloy

type GetGitlabWebhookInput struct {
	payload string
}

func (uc *Usecase) GetGitlabWebhook(input *GetGitlabWebhookInput) error {
	_, err := uc.webhook.ParseGitlabWebhook(input.payload)
	uc.queue.EnQueue(nil)

	return err
}
