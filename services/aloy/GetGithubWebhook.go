package aloy

type GetGithubWebhookInput struct {
	payload string
}

func (uc *Usecase) GetGithubWebhook(input *GetGithubWebhookInput) error {
	_, err := uc.webhook.ParseGithubWebhook(input.payload)
	uc.queue.EnQueue(nil)

	return err
}
