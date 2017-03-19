package webhook

import "time"

// GitlabProjectWebhookResult struct
type GitlabProjectWebhookResult struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Homepage          string `json:"homepage"`
	PathWithNamespace string `json:"path_with_namespace"`
	HTTPURL           string `json:"http_url"`
}

// GitlabCommitsWebhookResult struct
type GitlabCommitsWebhookResult struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GitlabWebhookResult struct
type GitlabWebhookResult struct {
	Ref        string                        `json:"ref"`
	After      string                        `json:"after"`
	UserName   string                        `json:"user_name"`
	UserAvatar string                        `json:"user_avatar"`
	Project    *GitlabProjectWebhookResult   `json:"project"`
	Commits    []*GitlabCommitsWebhookResult `json:"commits"`
}

func (g *GitlabWebhookResult) mapToWebHook() *Webhook {
	var commitID, commitMessage, commitTimestamp string
	for _, commit := range g.Commits {
		if commit.ID == g.After {
			commitID = commit.ID
			commitMessage = commit.Message
			commitTimestamp = commit.Timestamp
		}
	}

	t, _ := time.Parse(time.RFC3339, commitTimestamp)
	t.Unix()

	return &Webhook{
		Provider:        "gitlab",
		SenderName:      g.UserName,
		SenderAvatar:    g.UserAvatar,
		CommitID:        commitID,
		CommitMessage:   commitMessage,
		CommitTimestamp: t.Unix(),
		Ref:             g.Ref,
		RepoName:        g.Project.PathWithNamespace,
		RepoHomepage:    g.Project.Homepage,
		RepoDesciption:  g.Project.Description,
		HTTPURL:         g.Project.HTTPURL,
	}
}
