package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
)

type GetAllCommitsByNameReq struct {
	Name   string
	Limit  int
	Offset int
}
type GetAllCommitsByNameRes struct {
	Commits []*entity.Commit
}

func (cs *CoverageService) GetAllCommitsByName(req *GetAllCommitsByNameReq) (*GetAllCommitsByNameRes, error) {
	commits, err := cs.store.GetAllCommitsByName(req.Name, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &GetAllCommitsByNameRes{commits}, nil
}
