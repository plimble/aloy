package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
)

type GetAllRepositorysReq struct {
	Limit  int
	Offset int
}
type GetAllRepositorysRes struct {
	Repositorys []*entity.Repository
}

func (cs *CoverageService) GetAllRepositorys(req *GetAllRepositorysReq) (*GetAllRepositorysRes, error) {
	repos, err := cs.store.GetAllRepositorys(req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &GetAllRepositorysRes{repos}, nil
}
