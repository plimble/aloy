package webhook

import (
	"encoding/json"

	"github.com/plimble/errors"
	"github.com/plimble/validator"
)

type ServiceInterface interface {
	ParseGithubWebhook(payload string) (*Webhook, error)
	ParseGitlabWebhook(payload string) (*Webhook, error)
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) ParseGithubWebhook(payload string) (*Webhook, error) {
	v := validator.New()
	v.RequiredString(payload, "payload")
	if err := v.GetError(); err != nil {
		return nil, errors.BadRequestErr(err)
	}

	gwh := &GithubWebHookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.MapToWebHook(), errors.WithStack(err)
}

func (s *Service) ParseGitlabWebhook(payload string) (*Webhook, error) {
	v := validator.New()
	v.RequiredString(payload, "payload")
	if err := v.GetError(); err != nil {
		return nil, errors.BadRequestErr(err)
	}

	gwh := &GitlabWebhookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.MapToWebHook(), errors.WithStack(err)
}
