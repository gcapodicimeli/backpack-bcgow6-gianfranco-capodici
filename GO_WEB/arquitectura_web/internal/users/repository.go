package users

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

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return us, nil
}

func (r *repository) Create(id int, name, lastName, email string, age int, heigth float64, active bool, date string) (u User, err error) {
	u = User{id, name, lastName, email, age, heigth, active, date}
	us = append(us, u)
	lastID = u.Id
	return
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
