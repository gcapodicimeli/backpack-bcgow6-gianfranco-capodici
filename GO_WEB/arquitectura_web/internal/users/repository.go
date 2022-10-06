package users

import "github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/pkg/store"

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
	Create(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (User, error)
	LastID() (int, error)
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

func (r *repository) Create(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (u User, err error) {
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
