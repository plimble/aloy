package aloy

type UsecaseInterface interface {
}

type Usecase struct {
	*Services
}

func New(services *Services) *Usecase {
	uc := &Usecase{
		Services: services,
	}

	return uc
}
