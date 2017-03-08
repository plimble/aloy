package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
	mockIdgen "github.com/plimble/aloy/services/coverage/idgen/mock"
	mockStore "github.com/plimble/aloy/services/coverage/store/mock"
	"github.com/plimble/errors"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type CreateCommitSuite struct {
	suite.Suite
	service Service
	store   *mockStore.Store
	idgen   *mockIdgen.IDGenerator
}

func TestCreateCommitSuite(t *testing.T) {
	suite.Run(t, &CreateCommitSuite{})
}

func (t *CreateCommitSuite) SetupTest() {
	t.store = &mockStore.Store{}
	t.idgen = &mockIdgen.IDGenerator{}

	t.service = &CoverageService{
		store: t.store,
		idgen: t.idgen,
	}
}

func (t *CreateCommitSuite) TestCreateCommit() {
	repository := &entity.Repository{
		Id:          "1",
		Name:        "repo_test",
		OwnerName:   "master_john",
		Source:      "github",
		Description: "repo_description",
		HomePage:    "http://test.com",
		CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
	}

	commit := &entity.Commit{
		Id:           "c_1",
		RepositoryId: "1",
		Ref:          "develop",
		SenderName:   "john",
		SenderAvatar: "http://gravata.com/john",
		CreatedAt:    time.Now().Truncate(time.Second).Format(time.RFC3339),
	}

	req := &CreateCommitReq{
		Commit:                "c_1",
		Ref:                   "develop",
		RepositoryOwner:       "master_john",
		RepositorySource:      "github",
		RepositoryName:        "repo_test",
		RepositoryDescription: "repo_description",
		RepositoryURL:         "http://test.com",
		SenderName:            "john",
		SenderAvatar:          "http://gravata.com/john",
	}
	expRes := &CreateCommitRes{
		Commit: commit,
	}

	t.store.On("GetRepository", "repo_test", "master_john", "github").Once().Return(repository, nil)
	t.store.On("CreateCommit", commit).Once().Return(nil)

	res, err := t.service.CreateCommit(req)
	t.store.AssertExpectations(t.T())

	t.NoError(err)
	t.Equal(expRes, res)
}

func (t *CreateCommitSuite) TestCreateCommitWithNewRepository() {
	repository := &entity.Repository{
		Id:          "1",
		Name:        "repo_test",
		OwnerName:   "master_john",
		Source:      "github",
		Description: "repo_description",
		HomePage:    "http://test.com",
		CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
	}

	commit := &entity.Commit{
		Id:           "c_1",
		RepositoryId: "1",
		Ref:          "develop",
		SenderName:   "john",
		SenderAvatar: "http://gravata.com/john",
		CreatedAt:    time.Now().Truncate(time.Second).Format(time.RFC3339),
	}

	req := &CreateCommitReq{
		Commit:                "c_1",
		Ref:                   "develop",
		RepositoryOwner:       "master_john",
		RepositorySource:      "github",
		RepositoryName:        "repo_test",
		RepositoryDescription: "repo_description",
		RepositoryURL:         "http://test.com",
		SenderName:            "john",
		SenderAvatar:          "http://gravata.com/john",
	}
	expRes := &CreateCommitRes{
		Commit: commit,
	}

	t.store.On("GetRepository", "repo_test", "master_john", "github").Once().Return(nil, errors.NotFound("test not found"))
	t.idgen.On("Generate").Once().Return("1")
	t.store.On("CreateRepository", repository).Once().Return(nil)
	t.store.On("CreateCommit", commit).Once().Return(nil)

	res, err := t.service.CreateCommit(req)
	t.store.AssertExpectations(t.T())

	t.NoError(err)
	t.Equal(expRes, res)
}
