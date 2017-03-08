package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
)

type GetAllCommitsByRepositoryAndRefReq struct {
	RepositoryId string
	Ref          string
	Limit        int
	Offset       int
}
type GetAllCommitsByRepositoryAndRefRes struct {
	Commits []*entity.Commit
}

func (cs *CoverageService) GetAllCommitsByRepositoryAndRef(req *GetAllCommitsByRepositoryAndRefReq) (*GetAllCommitsByRepositoryAndRefRes, error) {
	commits, err := cs.store.GetAllCommitsByRepositoryAndRef(req.RepositoryId, req.Ref, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &GetAllCommitsByRepositoryAndRefRes{commits}, nil
}
