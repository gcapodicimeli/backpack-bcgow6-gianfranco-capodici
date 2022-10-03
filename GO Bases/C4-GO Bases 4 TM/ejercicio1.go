package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	message string
}

func (e *myCustomError) Error() string {
	return e.message
}

func myFunctionTest(salary int) error {
	if salary < 150_000 {
		return &myCustomError{
			message: "error: el salario ingresado no alcanza el minimo imponible",
		}
	}

	return nil
}

func main() {
	salary := 20_000

	err := myFunctionTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}
