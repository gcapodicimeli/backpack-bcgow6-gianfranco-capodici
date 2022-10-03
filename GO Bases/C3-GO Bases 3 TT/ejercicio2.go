package main

import "fmt"

type Product struct {
	name     string
	price    float64
	quantity int
}

type User struct {
	name     string
	lastName string
	email    string
	products []Product
}

func newProduct(name string, price float64) (newP Product) {
	newP.name = name
	newP.price = price
	return
}

func addProduct(u *User, p *Product, quantity int) {
	u.products = append(u.products, *p)
	p.quantity += quantity
}

func deleteAllProducts(u *User) {
	u.products = []Product{}
}

func main() {
	var u User
	u.name = "John"
	u.lastName = "Doe"
	u.email = "john.doe@mail.com"
	p := newProduct("Mouse", 5000.5)
	p1 := newProduct("Monitor", 5400.5)

	fmt.Printf("User: %v\n", u)

	addProduct(&u, &p, 5)
	addProduct(&u, &p1, 2)

	fmt.Printf("User: %v\n", u)

	deleteAllProducts(&u)

	fmt.Printf("User: %v\n", u)
}
