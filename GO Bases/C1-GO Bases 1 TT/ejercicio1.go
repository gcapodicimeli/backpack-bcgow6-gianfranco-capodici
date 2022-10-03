package main

import "fmt"

func main() {
	name := []string{"g", "i", "a", "n"}

	fmt.Println("Longitud: ", len(name))

	for _, letra := range name {
		fmt.Println(letra)
	}
}
