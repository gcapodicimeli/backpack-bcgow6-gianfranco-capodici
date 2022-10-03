// !- incompleto
package main

import "fmt"

type store struct {
	products []product
}

type product struct {
	typeProduct string
	name        string
	price       float64
}

type Product interface {
	CalculateCost() float64
}

type Ecommerce interface {
	Total()
	Add(p Product)
}

func (p product) CalculateCost() (newPrice float64) {
	switch p.typeProduct {
	case "pequeño":
		newPrice = p.price
	case "mediano":
		newPrice = p.price + p.price*0.03
	case "grande":
		newPrice = p.price + p.price*0.06 + 2500
		// ? controlar el error
	}
	return
}

func (s store) Total() {

}

func (e *store) Add(p product) {

}

func newProduct(typeProduct, name string, price float64) Product {

	return product{typeProduct: typeProduct, name: name, price: price}
}

func newStore() (ecommerceRes Ecommerce) {

	return
}

func main() {
	p := newProduct("grande", "mouse", 5000.3)
	e := newStore()
	e.Add(p)

	fmt.Printf("El costo final será de: %.2f\n", p.CalculateCost())
}
