package main

import (
	"errors"
	"fmt"
	"os"
)

// * inciso a)
func calculateSalary(hours, pricePerHour int) (salary int, err error) {
	if hours < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
		return
	}

	salary = hours * pricePerHour

	if salary >= 150_000 {
		salary = salary - salary/10
	}

	return
}

func main() {
	salary, err := calculateSalary(150, 10000)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("El salario mensual es $%d\n", salary)
}
