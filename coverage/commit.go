package coverage

// Commit data
type Commit struct {
	ID           string
	Repo         string
	Provider     string
	SenderName   string
	SenderAvatar string
	Message      string
	Timestamp    int64
	Ref          string
	Status       Status
	Cov          float64
}
