package users

import (
	"fmt"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/pkg/store"
)

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

var us []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (User, error)
	LastID() (int, error)
	Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (User, error)
	Delete(id int) error
	UpdateName(id int, name string) (User, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]User, error) {
	var us []User
	r.db.Read(&us)
	return us, nil
}

func (r *repository) Store(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (u User, err error) {
	r.db.Read(&us)
	u = User{id, name, lastName, email, age, heigth, active, date}
	us = append(us, u)

	if err := r.db.Write(us); err != nil {
		return User{}, err
	}

	lastID = u.Id
	return
}

func (r *repository) LastID() (int, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return 0, err
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}

func (r *repository) Update(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (user User, err error) {
	user = User{Name: name, LastName: lastName, Email: email, Age: age, Height: heigth, Active: active, Date: date}
	r.db.Read(&us)
	for i := range us {
		if us[i].Id == id {
			user.Id = id
			us[i] = user
			if err := r.db.Write(us); err != nil {
				return User{}, err
			}

			return
		}
	}

	user = User{}
	err = fmt.Errorf("ID %d no encontrado", id)
	return
}

func (r *repository) Delete(id int) (err error) {
	r.db.Read(&us)
	for i := range us {
		if us[i].Id == id {
			us = append(us[:i], us[i+1:]...)
			r.db.Write(us)
			return
		}
	}

	err = fmt.Errorf("ID %d no encontrado", id)
	return
}

func (r *repository) UpdateName(id int, name string) (u User, err error) {
	r.db.Read(&us)
	for i := range us {
		if us[i].Id == id {
			us[i].Name = name
			r.db.Write(us)
			u = us[i]
			return
		}
	}

	err = fmt.Errorf("ID %d no encontrado", id)
	return
}
