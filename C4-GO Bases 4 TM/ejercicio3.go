package main

import (
	"fmt"
	"os"
)

func myFunctionTest(salary int) {
	if salary < 150_000 {
		err := fmt.Errorf("error: el minimo imponible es de 150000 y el salario ingresado es de %d", salary)
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

func main() {
	myFunctionTest(20_000)
}
