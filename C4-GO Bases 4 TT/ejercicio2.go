package main

import (
	"errors"
	"fmt"
	"os"
)

type Client struct {
	File        int
	fullName    string
	Dni         int
	numberPhone int
	direction   string
}

// * Tarea 1
func generateID(id int) (client Client, err error) {
	if id == 0 {
		err = fmt.Errorf("ID no puede ser 0")
		return
	}

	client.File = id
	client.fullName = "Gianfranco Capodici"
	client.Dni = 39784345
	client.numberPhone = 12312323
	client.direction = "Direccion del cliente"
	return
}

// * Tarea 2
func readFile(fileName string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile(fileName)

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
}

// * Tarea 3
func validateFields(client Client) (isEmpty bool) {
	if client.File == 0 || client.fullName == "" || client.Dni == 0 || client.numberPhone == 0 || client.direction == "" {
		isEmpty = false

	} else {
		isEmpty = true
	}

	return
}

func main() {
	// --- Tarea 1 ---
	client, err := generateID(1)
	if err != nil {
		panic(err)
	}
	// ---------------

	// --- Tarea 2 ---
	readFile("customers2.txt")
	// ---------------

	// --- Tarea 3 ---
	res := validateFields(client)
	if !res {
		err := errors.New("error: los campos no pueden estar vacio")
		fmt.Print(err)
	}
	// ---------------

	fmt.Println("\nFin de la ejecución")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	fmt.Println("No han quedado archivos abiertos")

	// res := fmt.Sprintf("Legajo: %d\nNombre Completo: %s\nDNI: %d\nNumero de telefono: %d\nDomicilio: %s",
	// 	client.File,
	// 	client.fullName,
	// 	client.Dni,
	// 	client.numberPhone,
	// 	client.direction,
	// )
	// os.WriteFile("./customers.txt", []byte(res), 0644)
}
