package webhook

import "time"

type GithubSenderWebhookResult struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type GithubHeadCommitWebHookResult struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type GithubRepositoryWebhookResult struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CloneURL    string `json:"clone_url"`
}

type GithubWebHookResult struct {
	Ref        string                         `json:"ref"`
	After      string                         `json:"after"`
	Repository *GithubRepositoryWebhookResult `json:"repository"`
	Sender     *GithubSenderWebhookResult     `json:"sender"`
	HeadCommit *GithubHeadCommitWebHookResult `json:"head_commit"`
}

func (g *GithubWebHookResult) mapToWebHook() *Webhook {
	t, _ := time.Parse(time.RFC3339, g.HeadCommit.Timestamp)

	return &Webhook{
		Provider:        "github",
		SenderName:      g.Sender.Login,
		SenderAvatar:    g.Sender.AvatarURL,
		CommitID:        g.HeadCommit.ID,
		CommitMessage:   g.HeadCommit.Message,
		CommitTimestamp: t.Unix(),
		Ref:             g.Ref,
		RepoName:        g.Repository.FullName,
		RepoHomepage:    g.Repository.URL,
		RepoDesciption:  g.Repository.Description,
		HTTPURL:         g.Repository.CloneURL,
	}
}
