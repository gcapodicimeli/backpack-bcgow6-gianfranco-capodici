package main

import "fmt"

func calcImpuesto(sueldo float64) float64 {
	if sueldo > 150_000 {
		return sueldo * 0.27
	} else if sueldo > 50_000 {
		return sueldo * 0.10
	}

	return 0
}

func main() {
	fmt.Println("El impuesto será de $", calcImpuesto(52000))
}
