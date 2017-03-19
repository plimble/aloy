package runner

// Status type for runner
type Status int

const (
	// PENDING wait for runner start
	PENDING Status = iota
	// SUCCESS runner run successful
	SUCCESS
	// FAILED has an error on running
	FAILED
	// UNKNOWN the runner error
	UNKNOWN
)

// Result of runner
type Result struct {
	Message Message
	Status  Status
	Cov     float64
	HTML    string
}
