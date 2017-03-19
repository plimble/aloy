package runner

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RunnerServiceSuite struct {
	suite.Suite
	service *service
}

func TestRunnerServiceSuite(t *testing.T) {
	suite.Run(t, &RunnerServiceSuite{})
}

func (t *RunnerServiceSuite) SetupTest() {
	t.service = NewService(Options{
		MaxQueue:          10,
		MaxRunner:         5,
		GithubAccessToken: "GithubUsername",
		GitlabAccessToken: "GithubPassword",
	}).(*service)
}

func (t *RunnerServiceSuite) TearDownTest() {
	t.service.Close()
}

func (t *RunnerServiceSuite) TestSettingOptions() {
	t.Equal(10, t.service.opt.MaxQueue)
	t.Equal(5, t.service.opt.MaxRunner)
	t.Nil(t.service.opt.ResultFunc)
	t.Nil(t.service.opt.RunnerFunc)
}

func (t *RunnerServiceSuite) TestRun() {
	quite := make(chan bool)
	t.service.SetRunnerFunc(func(msg Message, result chan Result, opt *Options) {
		result <- Result{
			Message: msg,
			Status:  PENDING,
		}

		result <- Result{
			Message: msg,
			Status:  SUCCESS,
		}

		result <- Result{
			Message: msg,
			Status:  FAILED,
		}
	})

	t.service.SetResultFunc(func(result Result) {
		t.Equal("1", result.Message.CommitID)
		if result.Status == FAILED {
			quite <- true
		}
	})

	t.service.Enqueue(Message{
		Provider:        "",
		SenderName:      "",
		SenderAvatar:    "",
		CommitID:        "1",
		CommitMessage:   "2",
		CommitTimestamp: 3,
		Ref:             "",
		RepoName:        "",
		RepoHomepage:    "",
		RepoDesciption:  "",
		HTTPURL:         "",
	})

	<-quite
}
