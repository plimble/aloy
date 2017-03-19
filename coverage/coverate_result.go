package coverage

type Status int

const (
	PENDING Status = iota
	SUCCESS
	FAILED
	UNKNOWN
)

type CoverateResult struct {
	Repo     string
	CommitID string
	Status   Status
	Cov      float64
	HTML     string
}
