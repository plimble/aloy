package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/plimble/aloy/services/coverage/entity"
	"github.com/plimble/errors"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type StoreSuite struct {
	suite.Suite
	store *Store
}

func TestStoreSuite(t *testing.T) {
	suite.Run(t, &StoreSuite{})
}

func (t *StoreSuite) SetupTest() {
	db, err := sql.Open("mysql", "root:root@/aloy_test")
	if err != nil {
		panic(err)
	}

	t.store = NewStore(db)
}

func (t *StoreSuite) TearDownTest() {
	t.store.db.Exec(`TRUNCATE TABLE repositorys`)
	t.store.db.Exec(`TRUNCATE TABLE commits`)
}

func (t *StoreSuite) TearDownSuite() {
	t.store.db.Exec(`DROP SCHEMA public CASCADE`)
	t.store.db.Exec(`CREATE SCHEMA public`)
	t.store.db.Close()
}

func (t *StoreSuite) TestCreateAndGetRepository() {
	repository := &entity.Repository{
		Id:          "1",
		Name:        "repo_test",
		OwnerName:   "owner_test",
		Source:      "github",
		Description: "repo_description",
		HomePage:    "http://test.com",
		CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
	}

	err := t.store.CreateRepository(repository)
	t.NoError(err)

	resRepository, err := t.store.GetRepository("repo_test", "owner_test", "github")
	t.NoError(err)
	t.Equal(repository, resRepository)
}
