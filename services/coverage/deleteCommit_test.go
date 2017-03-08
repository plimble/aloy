package coverage

import (
	mockIdgen "github.com/plimble/aloy/services/coverage/idgen/mock"
	mockStore "github.com/plimble/aloy/services/coverage/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeleteCommitSuite struct {
	suite.Suite
	service Service
	store   *mockStore.Store
	idgen   *mockIdgen.IDGenerator
}

func TestDeleteCommitSuite(t *testing.T) {
	suite.Run(t, &DeleteCommitSuite{})
}

func (t *DeleteCommitSuite) SetupTest() {
	t.store = &mockStore.Store{}
	t.idgen = &mockIdgen.IDGenerator{}

	t.service = &CoverageService{
		store: t.store,
		idgen: t.idgen,
	}
}

func (t *DeleteCommitSuite) TestDeleteCommit() {
	req := &DeleteCommitReq{
		CommitId: "1",
	}
	expRes := &DeleteCommitRes{}

	t.store.On("DeleteCommit", "1").Once().Return(nil)

	res, err := t.service.DeleteCommit(req)
	t.store.AssertExpectations(t.T())
	t.NoError(err)
	t.Equal(expRes, res)
}
