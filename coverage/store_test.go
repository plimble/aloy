package coverage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LeveldbStoreSuite struct {
	suite.Suite
	store *leveldbstore
}

func TestLeveldbStoreSuite(t *testing.T) {
	suite.Run(t, &LeveldbStoreSuite{})
}

func (t *LeveldbStoreSuite) SetupSuite() {
	var err error
	store, err := NewLeveldbStore("./db_test")
	t.store = store.(*leveldbstore)
	if err != nil {
		t.Error(err)
	}
}

func (t *LeveldbStoreSuite) TearDownSuite() {
	t.store.Close()
	os.RemoveAll("./db_test")
}

func (t *LeveldbStoreSuite) TestSaveAndGetCommits() {
	var err error
	expCommits := []*Commit{
		{"xxxx1", "usr/repo", "github", "name", "avatar", "msg", 1000, "ref/head/master", SUCCESS, 100.00},
		{"xxxx2", "usr/repo", "github", "name", "avatar", "msg", 1001, "ref/head/master", SUCCESS, 100.00},
		{"xxxx3", "usr/repo", "github", "name", "avatar", "msg", 1002, "ref/head/master", SUCCESS, 100.00},
		{"xxxx4", "usr/repo", "github", "name", "avatar", "msg", 1003, "ref/head/master", SUCCESS, 100.00},
		{"xxxx5", "usr/repo", "github", "name", "avatar", "msg", 1004, "ref/head/staging", SUCCESS, 100.00},
		{"xxxx6", "usr/repo", "github", "name", "avatar", "msg", 1005, "ref/head/staging", SUCCESS, 100.00},
		{"xxxx7", "usr/repo", "github", "name", "avatar", "msg", 1006, "ref/head/staging", SUCCESS, 100.00},
		{"xxxx8", "usr/repo", "github", "name", "avatar", "msg", 1007, "ref/head/staging", SUCCESS, 100.00},
	}

	for _, commit := range expCommits {
		err = t.store.SaveCommit(commit)
		t.NoError(err)
	}

	commits, err := t.store.GetCommits("usr/repo", "ref/head/master", "", 10)
	t.NoError(err)
	t.Len(commits, 4)

	commits, err = t.store.GetCommits("", "", "usr/repo:ref/head/master:8997", 10)
	t.NoError(err)
	t.Len(commits, 2)
}

func (t *LeveldbStoreSuite) TestSaveAndGetCoverageResult() {
	expResult := &CoverateResult{
		Repo:     "usr/repo",
		CommitID: "xxxx1",
		Status:   SUCCESS,
		Cov:      100.00,
	}

	err := t.store.SaveCoverateResult(expResult)
	t.NoError(err)

	result, err := t.store.GetCoverateResult(expResult.Repo, expResult.CommitID)
	t.NoError(err)
	t.Equal(expResult, result)
}
