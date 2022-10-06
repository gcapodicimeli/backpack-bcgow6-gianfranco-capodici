package main

import (
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/cmd/server/handler"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/internal/users"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)

	u := handler.NewUser(service)

	router := gin.Default()

	ur := router.Group("/users")
	ur.POST("/", u.Create)
	ur.GET("/", u.GetAll)

	router.Run()
}
