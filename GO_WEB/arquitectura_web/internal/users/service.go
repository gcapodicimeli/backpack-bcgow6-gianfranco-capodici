package users

type Service interface {
	GetAll() ([]User, error)
	Create(name, lastName, email string, age int, heigth float64, active bool, date string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() (users []User, err error) {
	users, err = s.repository.GetAll()
	if err != nil {
		return
	}

	return
}

func (s *service) Create(name, lastName, email string, age int, heigth float64, active bool, date string) (user User, err error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return
	}
	lastID++

	user, err = s.repository.Create(lastID, name, lastName, email, age, heigth, active, date)
	if err != nil {
		return
	}

	return
}
