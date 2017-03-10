package testrunner

type Message struct {
	Provider       string
	SenderName     string
	SenderAvatar   string
	Commit         string
	Ref            string
	RepoName       string
	RepoHomepage   string
	RepoDesciption string
	HTTPURL        string
}

type Result struct {
	Cov  float64
	HTML string
}

type RunnerFunc func(msg Message, result chan Result, opt ServiceOptions)
type ResultFunc func(result Result)

type ServiceInterface interface {
	Enqueue(msg Message)
}

type ServiceOptions struct {
	MaxQueue       int
	MaxRunner      int
	GitlabUsername string
	GitLabPassword string
	GithubUsername string
	GithubPassword string
	RunnerFunc     RunnerFunc
	ResultFunc     ResultFunc
}

type Service struct {
	q       chan Message
	channel chan Message
	result  chan Result
	opt     ServiceOptions
}

func New(opt ServiceOptions) *Service {
	s := &Service{
		opt: opt,
	}

	if s.opt.MaxRunner == 0 {
		s.opt.MaxRunner = 5
	}

	if s.opt.MaxQueue == 0 {
		s.opt.MaxQueue = 100
	}

	if s.opt.RunnerFunc == nil {
		s.opt.RunnerFunc = DefaultRunnerFunc
	}

	s.q = make(chan Message, s.opt.MaxQueue)
	s.channel = make(chan Message, s.opt.MaxRunner)
	s.result = make(chan Result, s.opt.MaxRunner)

	for i := 0; i < s.opt.MaxRunner; i++ {
		runner := newRunner(s.channel, s.result, s.opt)
		runner.start()
	}

	go s.dispatch()

	return s
}

func (s *Service) Enqueue(msg Message) {
	s.q <- msg
}

func (s *Service) dispatch() {
	for {
		select {
		case msg := <-s.q:
			go func(msg Message) {
				s.channel <- msg
			}(msg)
		}
	}
}
