package testrunner

import (
	"fmt"
)

type Runner struct {
	channel chan Message
	result  chan Result
	quit    chan bool
	opt     ServiceOptions
}

func newRunner(channel chan Message, result chan Result, opt ServiceOptions) *Runner {
	return &Runner{
		channel: channel,
		result:  result,
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
				r.opt.RunnerFunc(msg, r.result, r.opt)
			case <-r.quit:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case result := <-r.result:
				if r.opt.ResultFunc != nil {
					r.opt.ResultFunc(result)
				}
			case <-r.quit:
				return
			}
		}
	}()
}
