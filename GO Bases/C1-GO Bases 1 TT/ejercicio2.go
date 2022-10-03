package main

import "fmt"

func main() {
	edad := 25
	isEmpleado := true
	antiguedad := 6
	sueldo := 200000

	if edad > 22 && isEmpleado && antiguedad > 5 {
		if sueldo > 100000 {
			fmt.Println("Si se puede realizar el préstamos sin intereses :)")
		} else {
			fmt.Println("Si se puede realizar el préstamos pero tendrá intereses")
		}
	} else {
		fmt.Println("No se puede realizar el préstamos")
	}
}
