package coverage

import (
	"github.com/plimble/aloy/services/coverage/gateway"
)

//go:generate mockery -name Service -case underscore -outpkg mock -output ./mock
type Service interface {
	CreateCommit(*CreateCommitReq) (*CreateCommitRes, error)

	DeleteRepositoryAndCommits(*DeleteRepositoryAndCommitsReq) (*DeleteRepositoryAndCommitsRes, error)
	DeleteCommit(*DeleteCommitReq) (*DeleteCommitRes, error)

	GetAllRepositorys(*GetAllRepositorysReq) (*GetAllRepositorysRes, error)
	GetAllCommitsByRepository(*GetAllCommitsByRepositoryReq) (*GetAllCommitsByRepositoryRes, error)
	GetAllCommitsByName(*GetAllCommitsByNameReq) (*GetAllCommitsByNameRes, error)
	GetAllCommitsByRepositoryAndRef(*GetAllCommitsByRepositoryAndRefReq) (*GetAllCommitsByRepositoryAndRefRes, error)
}

type CoverageService struct {
	idgen gateway.IDGenerator
	store gateway.Store
}

func NewService(idgen gateway.IDGenerator, store gateway.Store) Service {
	return &CoverageService{idgen, store}
}
