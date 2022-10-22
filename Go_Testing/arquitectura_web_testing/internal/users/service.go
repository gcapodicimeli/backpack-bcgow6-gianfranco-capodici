package users

import "github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/domain"

type Service interface {
	GetAll() ([]domain.User, error)
	Store(name, lastName, email string, age int, heigth float64, active bool, date string) (domain.User, error)
	Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (domain.User, error)
	Delete(id int) error
	UpdateName(id int, name string) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() (users []domain.User, err error) {
	users, err = s.repository.GetAll()
	if err != nil {
		return
	}

	return
}

func (s *service) Store(name, lastName, email string, age int, heigth float64, active bool, date string) (user domain.User, err error) {
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

func (s *service) Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (user domain.User, err error) {
	user, err = s.repository.Update(id, name, lastName, email, age, heigth, active, date)
	if err != nil {
		return
	}
	return
}

func (s *service) Delete(id int) (err error) {
	err = s.repository.Delete(id)
	if err != nil {
		return
	}
	return
}

func (s *service) UpdateName(id int, name string) (u domain.User, err error) {
	u, err = s.repository.UpdateName(id, name)
	if err != nil {
		return
	}
	return
}
