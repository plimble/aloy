package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
	mockIdgen "github.com/plimble/aloy/services/coverage/idgen/mock"
	mockStore "github.com/plimble/aloy/services/coverage/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type GetAllCommitsByRepositoryAndRefSuite struct {
	suite.Suite
	service Service
	store   *mockStore.Store
	idgen   *mockIdgen.IDGenerator
}

func TestGetAllCommitsByRepositoryAndRefSuite(t *testing.T) {
	suite.Run(t, &GetAllCommitsByRepositoryAndRefSuite{})
}

func (t *GetAllCommitsByRepositoryAndRefSuite) SetupTest() {
	t.store = &mockStore.Store{}
	t.idgen = &mockIdgen.IDGenerator{}

	t.service = &CoverageService{
		store: t.store,
		idgen: t.idgen,
	}
}

func (t *GetAllCommitsByRepositoryAndRefSuite) TestGetAllCommitsByRepositoryAndRef() {
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

	req := &GetAllCommitsByRepositoryAndRefReq{
		RepositoryId: "1",
		Ref:          "develop",
		Limit:        5,
		Offset:       0,
	}
	expRes := &GetAllCommitsByRepositoryAndRefRes{
		Commits: commits,
	}

	t.store.On("GetAllCommitsByRepositoryAndRef", "1", "develop", 5, 0).Once().Return(commits, nil)

	res, err := t.service.GetAllCommitsByRepositoryAndRef(req)
	t.store.AssertExpectations(t.T())
	t.NoError(err)
	t.Equal(expRes, res)
}
