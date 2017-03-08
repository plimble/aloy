package gateway

import (
	"github.com/plimble/aloy/services/coverage/entity"
)

//go:generate mockery -name Store -case underscore -outpkg mock -output ../store/mock
type Store interface {
	CreateRepository(repository *entity.Repository) error
	CreateCommit(commit *entity.Commit) error

	DeleteRepositoryAndCommits(repositoryId string) error
	DeleteCommit(commitId string) error

	GetRepository(repositoryName, repositoryOwnerName, repositorySource string) (*entity.Repository, error)
	GetAllRepositorys(limit, offset int) ([]*entity.Repository, error)
	GetAllCommitsByRepository(repositoryId string, limit, offset int) ([]*entity.Commit, error)
	GetAllCommitsByName(name string, limit, offset int) ([]*entity.Commit, error)
	GetAllCommitsByRepositoryAndRef(repositoryId, ref string, limit, offset int) ([]*entity.Commit, error)
}
