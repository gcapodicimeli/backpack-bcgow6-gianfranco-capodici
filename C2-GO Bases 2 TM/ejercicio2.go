package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...float64) (float64, error) {
	var totalNotas, cantNotas float64

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Ninguna nota puede ser negativa")
		} else {
			totalNotas += nota
			cantNotas += 1
		}
	}

	return totalNotas / cantNotas, nil
}

func main() {
	result, err := promedio(6, 7, 8)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("El promedio del alumno es:", result)
	}
}
