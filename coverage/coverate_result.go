package coverage

// Status coverate testing
type Status int

const (
	// PENDING coverate testing
	PENDING Status = iota
	// SUCCESS coverate testing
	SUCCESS
	// FAILED coverate testing
	FAILED
	// UNKNOWN error on coverate testing
	UNKNOWN
)

// CoverateResult data
type CoverateResult struct {
	Repo     string
	CommitID string
	Status   Status
	Cov      float64
	HTML     string
}
