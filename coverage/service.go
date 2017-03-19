package coverage

// Service interface
type Service interface {
	GetCoverateResult(repo, commitID string) (*CoverateResult, error)
	GetCommits(repo, ref, last string, limit int) ([]*Commit, error)
	SaveCommit(commit *Commit) error
	SaveCoverateResult(result *CoverateResult) error
}

type service struct {
	store Store
}

// NewService func
func NewService(store Store) Service {
	return &service{store}
}

func (s *service) GetCoverateResult(repo, commitID string) (*CoverateResult, error) {
	return s.store.GetCoverateResult(repo, commitID)
}

func (s *service) GetCommits(repo, ref, last string, limit int) ([]*Commit, error) {
	return s.store.GetCommits(repo, ref, last, limit)
}

func (s *service) SaveCommit(commit *Commit) error {
	return s.store.SaveCommit(commit)
}

func (s *service) SaveCoverateResult(result *CoverateResult) error {
	return s.store.SaveCoverateResult(result)
}
