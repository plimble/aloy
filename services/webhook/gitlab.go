package webhook

type GitlabProjectWebhookResult struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Homepage          string `json:"homepage"`
	PathWithNamespace string `json:"path_with_namespace"`
	HTTPURL           string `json:"http_url"`
}

type GitlabWebhookResult struct {
	Ref        string                      `json:"ref"`
	After      string                      `json:"after"`
	UserName   string                      `json:"user_name"`
	UserAvatar string                      `json:"user_avatar"`
	Project    *GitlabProjectWebhookResult `json:"project"`
}

func (g *GitlabWebhookResult) MapToWebHook() *Webhook {
	return &Webhook{
		SenderName:     g.UserName,
		SenderAvatar:   g.UserAvatar,
		Commit:         g.After,
		Ref:            g.Ref,
		RepoName:       g.Project.PathWithNamespace,
		RepoHomepage:   g.Project.Homepage,
		RepoDesciption: g.Project.Description,
		HTTPURL:        g.Project.HTTPURL,
	}
}
