package coverage

import (
	"github.com/plimble/aloy/services/coverage/entity"
	mockIdgen "github.com/plimble/aloy/services/coverage/idgen/mock"
	mockStore "github.com/plimble/aloy/services/coverage/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type GetAllRepositorysSuite struct {
	suite.Suite
	service Service
	store   *mockStore.Store
	idgen   *mockIdgen.IDGenerator
}

func TestGetAllRepositorysSuite(t *testing.T) {
	suite.Run(t, &GetAllRepositorysSuite{})
}

func (t *GetAllRepositorysSuite) SetupTest() {
	t.store = &mockStore.Store{}
	t.idgen = &mockIdgen.IDGenerator{}

	t.service = &CoverageService{
		store: t.store,
		idgen: t.idgen,
	}
}

func (t *GetAllRepositorysSuite) TestGetAllRepositorys() {
	repos := []*entity.Repository{
		{
			Id:          "1",
			Name:        "repo_test",
			Description: "repo_description",
			HomePage:    "http://test.com",
			CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
		},
		{
			Id:          "2",
			Name:        "repo_test_2",
			Description: "repo_description_2",
			HomePage:    "http://test.com_2",
			CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
		},
	}

	req := &GetAllRepositorysReq{
		Limit:  5,
		Offset: 0,
	}
	expRes := &GetAllRepositorysRes{
		Repositorys: repos,
	}

	t.store.On("GetAllRepositorys", 5, 0).Once().Return(repos, nil)

	res, err := t.service.GetAllRepositorys(req)
	t.store.AssertExpectations(t.T())
	t.NoError(err)
	t.Equal(expRes, res)
}
