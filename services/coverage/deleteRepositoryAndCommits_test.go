package coverage

import (
	mockIdgen "github.com/plimble/aloy/services/coverage/idgen/mock"
	mockStore "github.com/plimble/aloy/services/coverage/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeleteRepositoryAndCommitsSuite struct {
	suite.Suite
	service Service
	store   *mockStore.Store
	idgen   *mockIdgen.IDGenerator
}

func TestDeleteRepositoryAndCommitsSuite(t *testing.T) {
	suite.Run(t, &DeleteRepositoryAndCommitsSuite{})
}

func (t *DeleteRepositoryAndCommitsSuite) SetupTest() {
	t.store = &mockStore.Store{}
	t.idgen = &mockIdgen.IDGenerator{}

	t.service = &CoverageService{
		store: t.store,
		idgen: t.idgen,
	}
}

func (t *DeleteRepositoryAndCommitsSuite) TestDeleteRepositoryAndCommits() {
	req := &DeleteRepositoryAndCommitsReq{
		RepositoryId: "1",
	}
	expRes := &DeleteRepositoryAndCommitsRes{}

	t.store.On("DeleteRepositoryAndCommits", "1").Once().Return(nil)

	res, err := t.service.DeleteRepositoryAndCommits(req)
	t.store.AssertExpectations(t.T())
	t.NoError(err)
	t.Equal(expRes, res)
}
