package webhook

import (
	"encoding/json"

	"github.com/plimble/errors"
)

//go:generate mockery -name Service -case underscore -outpkg mock -output ./mock

// Service webhook interface
type Service interface {
	ParseGithubWebhook(payload []byte) (*Webhook, error)
	ParseGitlabWebhook(payload []byte) (*Webhook, error)
}

type service struct{}

// NewService webhook
func NewService() Service {
	return &service{}
}

func (s *service) ParseGithubWebhook(payload []byte) (*Webhook, error) {
	if payload == nil {
		return nil, errors.BadRequest("payload is missing")
	}

	gwh := &GithubWebHookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.mapToWebHook(), errors.WithStack(err)
}

func (s *service) ParseGitlabWebhook(payload []byte) (*Webhook, error) {
	if payload == nil {
		return nil, errors.BadRequest("payload is missing")
	}

	gwh := &GitlabWebhookResult{}
	err := json.Unmarshal([]byte(payload), gwh)

	return gwh.mapToWebHook(), errors.WithStack(err)
}
