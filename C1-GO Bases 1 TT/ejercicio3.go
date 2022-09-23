package main

import "fmt"

func main() {
	number := 9
	months := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	varMap := months[number]
	fmt.Println(number, ", "+varMap)

	switch number {
	case 1:
		fmt.Println(number, ", Enero")
	case 2:
		fmt.Println(number, ", Febrero")
	case 3:
		fmt.Println(number, ", Marzo")
	case 4:
		fmt.Println(number, ", Abril")
	case 5:
		fmt.Println(number, ", Mayo")
	case 6:
		fmt.Println(number, ", Junio")
	case 7:
		fmt.Println(number, ", Julio")
	case 8:
		fmt.Println(number, ", Agosto")
	case 9:
		fmt.Println(number, ", Septiembre")
	case 10:
		fmt.Println(number, ", Octubre")
	case 11:
		fmt.Println(number, ", Noviembre")
	case 12:
		fmt.Println(number, ", Diciembre")
	default:
		fmt.Println("No existe ese mes")
	}
}
