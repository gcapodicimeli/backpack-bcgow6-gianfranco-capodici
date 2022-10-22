package users

import (
	"fmt"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/domain"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/pkg/store"
)

var us []domain.User

type Repository interface {
	GetAll() ([]domain.User, error)
	Store(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (domain.User, error)
	LastID() (int, error)
	Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (domain.User, error)
	Delete(id int) error
	UpdateName(id int, name string) (domain.User, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.User, error) {
	var us []domain.User
	err := r.db.Read(&us)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (r *repository) Store(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (u domain.User, err error) {
	err = r.db.Read(&us)
	if err != nil {
		return
	}
	u = domain.User{Id: id, Name: name, LastName: lastName, Email: email, Age: age, Height: heigth, Active: active, Date: date}
	us = append(us, u)

	if err := r.db.Write(us); err != nil {
		return domain.User{}, err
	}

	// lastID = u.Id
	return
}

func (r *repository) LastID() (int, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return 0, err
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}

func (r *repository) Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (user domain.User, err error) {
	user = domain.User{Name: name, LastName: lastName, Email: email, Age: age, Height: heigth, Active: active, Date: date}
	err = r.db.Read(&us)
	if err != nil {
		return
	}
	for i := range us {
		if us[i].Id == id {
			user.Id = id
			us[i] = user
			if err := r.db.Write(us); err != nil {
				return domain.User{}, err
			}

			return
		}
	}

	user = domain.User{}
	err = fmt.Errorf("ID %d no encontrado", id)
	return
}

func (r *repository) Delete(id int) (err error) {
	err = r.db.Read(&us)
	if err != nil {
		return
	}
	for i := range us {
		if us[i].Id == id {
			us = append(us[:i], us[i+1:]...)
			err = r.db.Write(us)
			if err != nil {
				return
			}
			return
		}
	}

	err = fmt.Errorf("ID %d no encontrado", id)
	return
}

func (r *repository) UpdateName(id int, name string) (u domain.User, err error) {
	err = r.db.Read(&us)
	if err != nil {
		return
	}
	for i := range us {
		if us[i].Id == id {
			us[i].Name = name
			err = r.db.Write(us)
			if err != nil {
				return
			}
			u = us[i]
			return
		}
	}

	err = fmt.Errorf("ID %d no encontrado", id)
	return
}
