package runner

// Service runner interface
type Service interface {
	Enqueue(msg Message)
	SetResultFunc(resultFunc ResultFunc)
	SetRunFunc(runFunc RunFunc)
	Close()
}

// Options runner
type Options struct {
	MaxQueue          int
	MaxRunner         int
	GitlabAccessToken string
	GithubAccessToken string
	RunFunc           RunFunc
	ResultFunc        ResultFunc
	GoTestTags        string
}

type service struct {
	queue      chan Message
	channel    chan Message
	quiteQueue chan bool
	opt        Options
	runner     []*runner
}

// NewService create a new service
func NewService(opt Options) Service {
	s := &service{
		opt: opt,
	}

	if s.opt.MaxRunner == 0 {
		s.opt.MaxRunner = 5
	}

	if s.opt.MaxQueue == 0 {
		s.opt.MaxQueue = 100
	}

	s.queue = make(chan Message, s.opt.MaxQueue)
	s.channel = make(chan Message, s.opt.MaxRunner)
	s.quiteQueue = make(chan bool)

	s.runner = make([]*runner, s.opt.MaxRunner)

	for i := 0; i < s.opt.MaxRunner; i++ {
		s.runner[i] = newRunner(i, s.channel, &s.opt)
		s.runner[i].run()
	}

	go s.dispatch()

	return s
}

func (s *service) Enqueue(msg Message) {
	s.queue <- msg
}

func (s *service) SetResultFunc(resultFunc ResultFunc) {
	s.opt.ResultFunc = resultFunc
}

func (s *service) SetRunFunc(runFunc RunFunc) {
	s.opt.RunFunc = runFunc
}

func (s *service) Close() {
	s.quiteQueue <- true

	for i := 0; i < s.opt.MaxRunner; i++ {
		s.runner[i].close()
	}

	close(s.channel)
	close(s.queue)
}

func (s *service) dispatch() {
	for {
		select {
		case msg := <-s.queue:
			go func(msg Message) {
				s.channel <- msg
			}(msg)
		case <-s.quiteQueue:
			return
		}
	}
}
