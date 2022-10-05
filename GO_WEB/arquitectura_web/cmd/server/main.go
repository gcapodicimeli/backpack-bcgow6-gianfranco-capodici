package main

import (
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/cmd/server/handler"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)

	u := handler.NewUser(service)

	router := gin.Default()

	ur := router.Group("/users")
	ur.POST("/", u.Create)
	ur.GET("/", u.GetAll)

	router.Run()
}
