package main

import (
	"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	Id       int
	Name     string
	LastName string
	Email    string
	Age      int
	Height   float64
	Activo   bool
	Date     string
}

func jsonToStruct(u *[]User) {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &u); err != nil {
		panic(err)
	}
}

func GetAll(ctx *gin.Context) {
	var u []User
	jsonToStruct(&u)
	ctx.JSON(http.StatusOK, u)
}

func GetUser(ctx *gin.Context) {
	var u []User
	jsonToStruct(&u)
	for _, user := range u {
		if strconv.Itoa(user.Id) == ctx.Param("id") {
			ctx.String(http.StatusOK, "El usuario es %v\n", user)
			return
		}
	}
	ctx.String(http.StatusNotFound, "El usuario con id %v no existe", ctx.Param("id"))
}

func SearchUsers(ctx *gin.Context) {
	var u []User
	var usersRes []User
	jsonToStruct(&u)
	for _, user := range u {
		if strconv.Itoa(user.Age) == ctx.Query("age") {
			usersRes = append(usersRes, user)
		}
	}
	if len(usersRes) != 0 {
		ctx.String(http.StatusOK, "Los usuarios con la edad %v son: %v\n", ctx.Query("age"), usersRes)
	} else {
		ctx.String(http.StatusNotFound, "No existen usuario con edad %v\n", ctx.Query("age"))
	}
}

func main() {
	router := gin.Default()

	// Ejercicio 1 TM
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Gianfranco",
		})
	})

	// Ejercicio 2 TM
	//router.GET("/users", GetAll)

	// Ejercicio 1 TT
	router.GET("/users", SearchUsers)

	// Ejercicio 2 TT
	router.GET("/users/:id", GetUser)

	router.Run()
}
