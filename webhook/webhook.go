package webhook

type Webhook struct {
	Provider        string
	SenderName      string
	SenderAvatar    string
	CommitID        string
	CommitMessage   string
	CommitTimestamp int64
	Ref             string
	RepoName        string
	RepoHomepage    string
	RepoDesciption  string
	HTTPURL         string
}
