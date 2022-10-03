package main

import "fmt"

const (
	dog            = "dog"
	cat            = "cat"
	hamster        = "hamster"
	spider         = "spider"
	foodPerDog     = 10.0
	foodPerCat     = 5.0
	foodPerHamster = 250.0
	foodPerSpider  = 150.0
)

func foodDog(quantity float64) float64 {
	return quantity * foodPerDog
}

func foodCat(quantity float64) float64 {
	return quantity * foodPerCat
}

func foodHamster(quantity float64) float64 {
	return quantity * foodPerHamster
}

func foodSpider(quantity float64) float64 {
	return quantity * foodPerSpider
}

func Animal(animal string) (func(quantity float64) float64, error) {
	switch animal {
	case dog:
		return foodDog, nil
	case cat:
		return foodCat, nil
	case hamster:
		return foodHamster, nil
	case spider:
		return foodSpider, nil
	default:
		return nil, fmt.Errorf("unknown animal: %s", animal)
	}
}

func main() {
	var amount float64

	animalDog, err := Animal(dog)
	if err != nil {
		panic(err)
	}
	animalCat, err := Animal(cat)
	if err != nil {
		panic(err)
	}
	animalHamster, err := Animal(hamster)
	if err != nil {
		panic(err)
	}
	animalSpider, err := Animal(spider)
	if err != nil {
		panic(err)
	}

	amount += animalDog(5)
	amount += animalCat(5)

	hamsterGr := animalHamster(4) / 1000
	amount += hamsterGr
	spiderGr := animalSpider(7) / 1000
	amount += spiderGr

	fmt.Printf("La cantidad de alimento a comprar es de %.2fkg\n", amount)
}
