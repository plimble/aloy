package testrunner

type Status int

const (
	PENDING Status = iota
	SUCCESS
	FAILED
	UNKNOWN
)

type Result struct {
	Message Message
	Status  Status
	Cov     float64
	HTML    string
}
