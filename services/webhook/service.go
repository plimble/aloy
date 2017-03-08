package webhook

import (
	"encoding/json"

	"github.com/plimble/errors"
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
	if payload == "" {
		return nil, errors.BadRequest("payload is missing")
	}

	gwh := &GithubWebHookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.MapToWebHook(), errors.WithStack(err)
}

func (s *Service) ParseGitlabWebhook(payload string) (*Webhook, error) {
	if payload == "" {
		return nil, errors.BadRequest("payload is missing")
	}

	gwh := &GitlabWebhookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.MapToWebHook(), errors.WithStack(err)
}
