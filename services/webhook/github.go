package webhook

type GithubSenderWebhookResult struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
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
}

func (g *GithubWebHookResult) MapToWebHook() *Webhook {
	return &Webhook{
		SenderName:     g.Sender.Login,
		SenderAvatar:   g.Sender.AvatarURL,
		Commit:         g.After,
		Ref:            g.Ref,
		RepoName:       g.Repository.FullName,
		RepoHomepage:   g.Repository.URL,
		RepoDesciption: g.Repository.Description,
		HTTPURL:        g.Repository.CloneURL,
	}
}
