package users

type Service interface {
	GetAll() ([]User, error)
	Create(name, lastName, email string, age int, heigth float64, active bool, date string)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	//TODO: hacer el get
	return
}

func (s *service) Create(name, lastName, email string, age int, heigth float64, active bool, date string) {
	//TODO: hacer el post
}
