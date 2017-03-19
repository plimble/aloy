package testrunner

type RunnerFunc func(msg Message, result chan Result, opt *Options)
type ResultFunc func(result Result)

type runner struct {
	id          int
	quiteRunner chan bool
	quiteResult chan bool
	result      chan Result
	channel     chan Message
	opt         *Options
}

func newRunner(id int, channel chan Message, opt *Options) *runner {
	return &runner{
		id,
		make(chan bool),
		make(chan bool),
		make(chan Result),
		channel,
		opt,
	}
}

func (r *runner) close() {
	r.quiteRunner <- true
	r.quiteResult <- true
}

func (r *runner) run() {
	go func() {
		for {
			select {
			case msg := <-r.channel:
				if r.opt.RunnerFunc != nil {
					r.opt.RunnerFunc(msg, r.result, r.opt)
				}
			case <-r.quiteRunner:
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
			case <-r.quiteResult:
				return
			}
		}
	}()
}
