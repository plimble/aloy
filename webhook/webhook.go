package webhook

type Webhook struct {
	Provider       string
	SenderName     string
	SenderAvatar   string
	Commit         string
	Ref            string
	RepoName       string
	RepoHomepage   string
	RepoDesciption string
	HTTPURL        string
}
