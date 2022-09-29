package main

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

// func operation(operacion string) func(valores ...int) int {

// }

func main() {

}
