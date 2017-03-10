package testrunner

import (
	"os/exec"
	"regexp"
	"strconv"
)

func DefaultRunnerFunc(msg Message, result chan Result, opt ServiceOptions) {
	var usr string
	var pwd string

	switch msg.Provider {
	case "github":
		usr = opt.GithubUsername
		pwd = opt.GithubPassword
	case "gitlab":
		usr = opt.GitlabUsername
		pwd = opt.GitLabPassword
	default:
		return
	}

	out, err := exec.Command("docker", "run", "--rm", "-a", "stdout", "-a", "stderr", msg.Provider, msg.RepoName, usr, pwd).CombinedOutput()
	if err != nil {
		return
	}
	content := string(out)

	re := regexp.MustCompile("-- cov:([0-9.]*) --")
	matches := re.FindStringSubmatch(content)
	if len(matches) == 2 {
		cov, _ := strconv.ParseFloat(matches[1], 64)
		result <- Result{
			Cov:  cov,
			HTML: content,
		}
	} else {
		result <- Result{
			Cov:  0,
			HTML: content,
		}
	}
}
