package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
)

type GetAllCommitsByRepositoryReq struct {
	RepositoryName string
	Limit          int
	Offset         int
}
type GetAllCommitsByRepositoryRes struct {
	Commits []*entity.Commit
}

func (cs *CoverageService) GetAllCommitsByRepository(req *GetAllCommitsByRepositoryReq) (*GetAllCommitsByRepositoryRes, error) {
	commits, err := cs.store.GetAllCommitsByRepository(req.RepositoryName, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &GetAllCommitsByRepositoryRes{commits}, nil
}
