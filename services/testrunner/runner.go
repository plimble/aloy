package testrunner

import (
	"fmt"
)

type Runner struct {
	channel    chan Message
	quit       chan bool
	runnerFunc RunnerFunc
	opt        ServiceOptions
}

func newRunner(channel chan Message, opt ServiceOptions) *Runner {
	return &Runner{
		channel: channel,
		quit:    make(chan bool),
		opt:     opt,
	}
}

func (r *Runner) start() {
	go func() {
		for {
			select {
			case msg := <-r.channel:
				fmt.Println("start runner", msg)
				r.opt.RunnerFunc(msg, r.opt)
			case <-r.quit:
				return
			}
		}
	}()
}
