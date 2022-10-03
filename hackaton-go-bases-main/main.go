package main

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

const (
	operation = "CREATE"
	ticketId  = 1
)

func transformStringToTicket(lines [][]string) (tickets []service.Ticket) {
	for _, line := range lines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		price, err := strconv.Atoi(line[5])
		if err != nil {
			panic(err)
		}

		t := service.Ticket{
			id,
			line[1],
			line[2],
			line[3],
			line[4],
			price,
		}
		tickets = append(tickets, t)
	}
	return
}

func main() {
	var tickets []service.Ticket

	// Funcion para obtener tickets del archivo csv
	f := &file.File{Path: "./tickets.csv"}
	lines, err := f.Read()
	if err != nil {
		panic(err)
	}

	// Funcion para convertir los string obtenidos del archivo en un []tickets
	tickets = transformStringToTicket(lines)
	serv := service.NewBookings(tickets)

	// * Funciones solicitadas
	switch operation {
	case "READ":
		tick, err := serv.Read(ticketId)
		if err != nil {
			panic(err)
		}
		fmt.Println(tick)
	case "CREATE":
		f.Write(tickets[1])
	case "UPDATE":

	case "DELETE":

	default:

	}
}
