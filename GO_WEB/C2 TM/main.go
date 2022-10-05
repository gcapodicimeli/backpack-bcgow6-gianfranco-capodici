package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

func NewUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "12345" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Para contemplar el caso que sea vacio
	if len(users) == 0 {
		req.Id = 1
	} else {
		req.Id = users[len(users)-1].Id + 1
	}
	users = append(users, req)
	c.JSON(http.StatusOK, req)
}

var users []User

func main() {
	router := gin.Default()

	// * Este Get lo hago para probar que funcione el Post
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Ejercicio 1
	router.POST("/users", NewUser)

	router.Run()
}
