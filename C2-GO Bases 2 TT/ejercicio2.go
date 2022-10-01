package main

import "fmt"

type Matrix struct {
	values      []int
	high        int
	width       int
	isQuadratic bool
	maxValue    int
}

func (m *Matrix) Set(high int, width int, values ...int) {
	m.values = values
	m.high = high
	m.width = width
	m.isQuadratic = high == width

	for i := 0; i < len(values); i++ {
		if values[i] > m.maxValue {
			m.maxValue = values[i]
		}
	}
}

func (m *Matrix) Print() {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.high; j++ {
			fmt.Print(m.values[i*m.high+j], "\t")
		}
		fmt.Print("\n")
	}

	fmt.Printf("\nEl valor mas alto es: %d\n", m.maxValue)
	if m.isQuadratic {
		fmt.Println("Es una matriz cuadrada")
	} else {
		fmt.Println("No es una matriz cuadrada")
	}
}

func main() {
	var m Matrix
	m.Set(2, 3, 9, 8, 5, 4, 26, 3)
	m.Print()
}
