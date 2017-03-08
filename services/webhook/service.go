package webhook

import (
	"encoding/json"

	"github.com/plimble/errors"
)

type ServiceInterface interface {
	ParseGithubWebhook(payload []byte) (*Webhook, error)
	ParseGitlabWebhook(payload []byte) (*Webhook, error)
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) ParseGithubWebhook(payload []byte) (*Webhook, error) {
	if payload == nil {
		return nil, errors.BadRequest("payload is missing")
	}

	gwh := &GithubWebHookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.MapToWebHook(), errors.WithStack(err)
}

func (s *Service) ParseGitlabWebhook(payload []byte) (*Webhook, error) {
	if payload == nil {
		return nil, errors.BadRequest("payload is missing")
	}

	gwh := &GitlabWebhookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.MapToWebHook(), errors.WithStack(err)
}
