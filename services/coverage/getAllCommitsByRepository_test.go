package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
	mockIdgen "github.com/plimble/aloy/services/coverage/idgen/mock"
	mockStore "github.com/plimble/aloy/services/coverage/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type GetAllCommitsByRepositorySuite struct {
	suite.Suite
	service Service
	store   *mockStore.Store
	idgen   *mockIdgen.IDGenerator
}

func TestGetAllCommitsByRepositorySuite(t *testing.T) {
	suite.Run(t, &GetAllCommitsByRepositorySuite{})
}

func (t *GetAllCommitsByRepositorySuite) SetupTest() {
	t.store = &mockStore.Store{}
	t.idgen = &mockIdgen.IDGenerator{}

	t.service = &CoverageService{
		store: t.store,
		idgen: t.idgen,
	}
}

func (t *GetAllCommitsByRepositorySuite) TestGetAllCommitsByRepository() {
	commits := []*entity.Commit{
		{
			Id:           "c_1",
			RepositoryId: "1",
			Ref:          "develop",
			SenderName:   "john",
			SenderAvatar: "http://gravata.com/john",
			CreatedAt:    time.Now().Truncate(time.Second).Format(time.RFC3339),
		},
		{
			Id:           "c_2",
			RepositoryId: "1",
			Ref:          "develop",
			SenderName:   "doe",
			SenderAvatar: "http://gravata.com/doe",
			CreatedAt:    time.Now().Truncate(time.Second).Format(time.RFC3339),
		},
	}

	repo := &entity.Repository{
		Id:          "1",
		Name:        "repo_test",
		Description: "repo_description",
		HomePage:    "http://test.com",
		CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
	}

	req := &GetAllCommitsByRepositoryReq{
		RepositoryName: "repo_test",
		Limit:          5,
		Offset:         0,
	}
	expRes := &GetAllCommitsByRepositoryRes{
		Commits: commits,
	}

	t.store.On("GetRepositoryByName", "repo_test").Once().Return(repo, nil)
	t.store.On("GetAllCommitsByRepository", "1", 5, 0).Once().Return(commits, nil)

	res, err := t.service.GetAllCommitsByRepository(req)
	t.store.AssertExpectations(t.T())
	t.NoError(err)
	t.Equal(expRes, res)
}
