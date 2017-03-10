package aloy

type UsecaseInterface interface {
	GetGithubWebhook(input *GetGithubWebhookInput) error
	GetGitlabWebhook(input *GetGitlabWebhookInput) error
}

type Usecase struct {
	*Services
}

func New(services *Services) *Usecase {
	uc := &Usecase{
		Services: services,
	}

	uc.testrunner.SetWaitResult(uc.SaveTestResult)

	return uc
}
