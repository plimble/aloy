package testrunner

import (
	"os/exec"
	"regexp"
	"strconv"
)

func DefaultRunnerFunc(msg Message, result chan Result, opt *Options) {
	var usr string
	var pwd string

	switch msg.Provider {
	case "github":
		usr = opt.GithubAccessToken
		pwd = "x-oauth-basic"
	case "gitlab":
		usr = "oauth2"
		pwd = opt.GitlabAccessToken
	default:
		return
	}

	result <- Result{
		Message: msg,
		Status:  PENDING,
	}

	out, err := exec.Command("docker", "run", "--rm", "-a", "stdout", "-a", "stderr", msg.Provider, msg.RepoName, usr, pwd, opt.GoTestTags).CombinedOutput()
	if err != nil {
		result <- Result{
			Message: msg,
			Status:  FAILED,
		}
	}
	content := string(out)

	re := regexp.MustCompile("-- cov:([0-9.]*) --")
	matches := re.FindStringSubmatch(content)
	if len(matches) == 2 {
		cov, _ := strconv.ParseFloat(matches[1], 64)
		result <- Result{
			Message: msg,
			Status:  SUCCESS,
			Cov:     cov,
			HTML:    content,
		}
	} else {
		result <- Result{
			Message: msg,
			Status:  UNKNOWN,
			HTML:    content,
		}
	}
}
