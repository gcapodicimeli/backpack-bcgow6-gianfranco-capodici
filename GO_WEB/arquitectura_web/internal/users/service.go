package users

type Service interface {
	GetAll() ([]User, error)
	Store(name, lastName, email string, age int, heigth float64, active bool, date string) (User, error)
	Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (User, error)
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

func (s *service) Store(name, lastName, email string, age int, heigth float64, active bool, date string) (user User, err error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return
	}
	lastID++

	user, err = s.repository.Store(lastID, name, lastName, email, age, heigth, active, date)
	if err != nil {
		return
	}

	return
}

func (s *service) Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (user User, err error) {
	user, err = s.repository.Update(id, name, lastName, email, age, heigth, active, date)
	if err != nil {
		return
	}
	return
}
