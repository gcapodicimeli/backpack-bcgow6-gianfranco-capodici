package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var cantEmployees int

	fmt.Println("La edad de Bemjamin es:", employees["Benjamin"])
	for _, employe := range employees {
		if employe > 21 {
			cantEmployees++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 años es:", cantEmployees)
	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)
}
