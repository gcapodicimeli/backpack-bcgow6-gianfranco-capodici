package main

import "fmt"

type Product struct {
	name     string
	price    float64
	quantity float64
}

type Service struct {
	name    string
	price   float64
	minutes float64
}

type Maintenance struct {
	name  string
	price float64
}

func addProducts(c chan float64, products ...Product) {
	var totalPrice float64
	for _, product := range products {
		totalPrice += product.price * product.quantity
	}
	c <- totalPrice
	fmt.Printf("El total de los productos es: $%.2f\n", totalPrice)
	close(c)
}

func addMaintenance(c chan float64, m ...Maintenance) {
	var totalMaintenance float64
	for _, maintenance := range m {
		totalMaintenance += maintenance.price
	}
	c <- totalMaintenance
	fmt.Printf("El total del mantenimiento es: $%.2f\n", totalMaintenance)
	close(c)
}

func main() {
	c := make(chan float64)
	cM := make(chan float64)

	p1 := Product{"Mouse", 5_000, 5}
	p2 := Product{"Monitor", 54_000, 2}
	p3 := Product{"Teclado", 6_000.5, 4}
	go addProducts(c, p1, p2, p3)

	m1 := Maintenance{"Mantenimiento1", 107.5}
	m2 := Maintenance{"Mantenimiento2", 130}
	go addMaintenance(cM, m1, m2)

	totalP := <-c
	totalM := <-cM

	fmt.Printf("El total de todo es: $%.2f\n", totalP+totalM)
}
