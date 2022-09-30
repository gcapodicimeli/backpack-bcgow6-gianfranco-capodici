package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(valores ...int) int {
	min := valores[0]
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}

	return min
}

func averageFunc(valores ...int) int {
	var totalValores, cantValores int

	for _, valor := range valores {
		if valor < 0 {
			//return 0, errors.New("Ninguna nota puede ser negativa")
		} else {
			totalValores += valor
			cantValores++
		}
	}

	return totalValores / cantValores //, nil
}

func maxFunc(valores ...int) int {
	var max int
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}

	return max
}

func operation(op string) (func(valores ...int) int, error) {
	switch op {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("No existe la operacion " + op)
	}
}

func main() {
	min, err := operation(minimum)
	if err != nil {
		panic(err)
	}
	average, err := operation(average)
	if err != nil {
		panic(err)
	}
	max, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minValue := min(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := average(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := max(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("minValue:", minValue)
	fmt.Println("averageValue:", averageValue)
	fmt.Println("maxValue:", maxValue)
}
