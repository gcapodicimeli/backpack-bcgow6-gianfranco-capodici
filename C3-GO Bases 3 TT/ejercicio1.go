package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func setName(u *User, newName, newLastName string) {
	u.name = newName
	u.lastName = newLastName
}

func setAge(u *User, newAge int) {
	u.age = newAge
}

func setEmail(u *User, newEmail string) {
	u.email = newEmail
}

func setPassword(u *User, newPassword string) {
	u.password = newPassword
}

func main() {
	user := User{
		name:     "John",
		lastName: "Doe",
		age:      18,
		email:    "john.doe@example.com",
		password: "123qwe",
	}

	fmt.Printf("Nombre: %s\nApellido: %s\nAge: %d\nCorreo: %s\nContraseña: %s\n",
		user.name, user.lastName, user.age, user.email, user.password,
	)

	setName(&user, "Juan", "Perez")
	setAge(&user, 28)
	setEmail(&user, "juan.peres@example.com")
	setPassword(&user, "nuevapass")

	fmt.Printf("Nombre: %s\nApellido: %s\nAge: %d\nCorreo: %s\nContraseña: %s\n",
		user.name, user.lastName, user.age, user.email, user.password,
	)
}
