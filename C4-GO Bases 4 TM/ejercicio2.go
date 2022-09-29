package main

import (
	"errors"
	"fmt"
	"os"
)

func myFunctionTest(salary int) error {
	if salary < 150_000 {
		return errors.New("error: el salario ingresado no alcanza el minimo imponible")
	}

	return nil
}

func main() {
	salary := 20000

	err := myFunctionTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}
