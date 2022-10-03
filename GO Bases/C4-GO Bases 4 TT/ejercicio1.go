package main

import (
	"fmt"
	"os"
)

// !- aclaracion de buenas practicas: todo los panic se manejan del main

func readFile(fileName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile(fileName)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	//fmt.Println(string(res))
}

func main() {
	readFile("customers2.txt")

	fmt.Println("ejecución finalizada")

}
