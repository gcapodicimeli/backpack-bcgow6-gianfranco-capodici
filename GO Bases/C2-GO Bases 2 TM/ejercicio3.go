package main

import "fmt"

func calcSalario(minutos float64, categoria string) float64 {
	horas := minutos / 60

	switch categoria {
	case "C":
		return horas * 1_000
	case "B":
		return (horas * 1_500) + (horas*1_500)*0.2
	case "A":
		return (horas * 3_000) + (horas*1_500)*0.5
	default:
		return 0
	}
}

func main() {
	fmt.Println("El salario es $", calcSalario(3000, "A"))
	fmt.Println("El salario es $", calcSalario(3000, "B"))
	fmt.Println("El salario es $", calcSalario(3000, "C"))
}
