package main

import (
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
}
