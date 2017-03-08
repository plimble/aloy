package queue

type ServiceInterface interface {
	EnQueue(msg *Message)
	DeQueue(msg *Message)
}

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) EnQueue(msg *Message) {

}

func (s *Service) DeQueue(msg *Message) {

}
