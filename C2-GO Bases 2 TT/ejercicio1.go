package main

import "fmt"

type alumno struct {
	Nombre   string
	Apellido string
	Dni      int
	Fecha    string
}

func (a alumno) showDetail() {
	fmt.Println("Nombre:", a.Nombre)
	fmt.Println("Apellido:", a.Apellido)
	fmt.Println("DNI:", a.Dni)
	fmt.Println("Fecha de ingreso:", a.Fecha)
}

func main() {
	al := alumno{
		Nombre:   "Juan",
		Apellido: "Perez",
		Dni:      1324242,
		Fecha:    "19-09-2022",
	}

	al.showDetail()
}
