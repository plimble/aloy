package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
	"github.com/plimble/errors"
)

type CreateCommitReq struct {
	SenderName            string
	SenderAvatar          string
	Commit                string
	Ref                   string
	RepositoryName        string
	RepositoryURL         string
	RepositoryDescription string
}
type CreateCommitRes struct {
	*entity.Commit
}

func (cs *CoverageService) CreateCommit(req *CreateCommitReq) (*CreateCommitRes, error) {
	repo, err := cs.store.GetRepositoryByName(req.RepositoryName)

	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	if err != nil && errors.IsNotFound(err) {
		repoId := cs.idgen.Generate()
		repo = entity.NewRepository(repoId, req.RepositoryName, req.RepositoryDescription, req.RepositoryURL)
		if err := cs.store.CreateRepository(repo); err != nil {
			return nil, err
		}
	}

	commit := entity.NewCommit(req.Commit, repo.Id, req.Ref, req.SenderName, req.SenderAvatar)
	if err := cs.store.CreateCommit(commit); err != nil {
		return nil, err
	}

	return &CreateCommitRes{commit}, nil
}
