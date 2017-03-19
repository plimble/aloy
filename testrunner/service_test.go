package testrunner

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestRunnerServiceSuite struct {
	suite.Suite
	service *service
}

func TestTestRunnerServiceSuite(t *testing.T) {
	suite.Run(t, &TestRunnerServiceSuite{})
}

func (t *TestRunnerServiceSuite) SetupTest() {
	t.service = NewService(Options{
		MaxQueue:       10,
		MaxRunner:      5,
		GithubUsername: "GithubUsername",
		GithubPassword: "GithubPassword",
		GitlabUsername: "GitlabUsername",
		GitLabPassword: "GitLabPassword",
	})
}

func (t *TestRunnerServiceSuite) TearDownTest() {
	t.service.Close()
}

func (t *TestRunnerServiceSuite) TestSettingOptions() {
	t.Equal(10, t.service.opt.MaxQueue)
	t.Equal(5, t.service.opt.MaxRunner)
	t.Nil(t.service.opt.ResultFunc)
	t.Nil(t.service.opt.RunnerFunc)
}

func (t *TestRunnerServiceSuite) TestRun() {
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
