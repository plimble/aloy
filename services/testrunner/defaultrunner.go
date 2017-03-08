package testrunner

import (
	"os/exec"
)

func defaultRunnerFunc(msg Message, opt ServiceOptions) {
	exec.Command("git", "clone")
}
