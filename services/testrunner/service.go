package testrunner

import "fmt"

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

const (
	PENDING = iota
	SUCCESS
	FAILED
	UNKNOWN
)

type Result struct {
	Status int
	Cov    float64
	HTML   string
}

type RunnerFunc func(msg Message, result chan Result, opt ServiceOptions)
type ResultFunc func(result Result)

type ServiceInterface interface {
	Enqueue(msg Message)
	SetWaitResult(resultFunc ResultFunc)
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
		s.startRunner()
		s.waitResult()
	}

	go s.dispatch()

	return s
}

func (s *Service) Enqueue(msg Message) {
	s.q <- msg
}

func (s *Service) SetWaitResult(resultFunc ResultFunc) {
	s.opt.ResultFunc = resultFunc
}

func (s *Service) waitResult() {
	go func() {
		for {
			select {
			case result := <-s.result:
				if s.opt.ResultFunc != nil {
					s.opt.ResultFunc(result)
				}
			}
		}
	}()
}

func (s *Service) startRunner() {
	go func() {
		for {
			select {
			case msg := <-s.channel:
				fmt.Println("start runner", msg)
				s.opt.RunnerFunc(msg, s.result, s.opt)
			}
		}
	}()
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
