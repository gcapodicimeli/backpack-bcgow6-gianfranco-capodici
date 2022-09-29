package main

import (
	"fmt"
	"os"
)

type product struct {
	id       int
	price    float64
	quantity int
}

func saveFile(products ...product) {
	res := fmt.Sprint("ID ; Precio ; Cantidad\n")

	for _, p := range products {
		res += fmt.Sprintf("%d ; %.2f ; %d\n", p.id, p.price, p.quantity)
	}

	err := os.WriteFile("./prueba.csv", []byte(res), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	product1 := product{
		id:       1,
		price:    100.0,
		quantity: 3,
	}

	product2 := product{
		id:       2,
		price:    150.5,
		quantity: 2,
	}

	product3 := product{
		id:       3,
		price:    600.9,
		quantity: 4,
	}

	saveFile(product1, product2, product3)
}
