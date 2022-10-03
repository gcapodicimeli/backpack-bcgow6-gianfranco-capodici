package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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

func jsonToStruct() {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	jsonData := []byte(`{"Id": 1, "Name": "Franco"}`)
	var u User
	if err := json.Unmarshal(jsonData, &u); err != nil {
		panic(err)
	}
	fmt.Println(u)
}

func GetAll(c *gin.Context) {
	//var u User
	jsonToStruct()

	// data, err := os.ReadFile("../users.json")
	// if err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(data, &u); err != nil {
	// 	panic(err)
	// }
	// c.Bind(&u) // ?- No sé que es, estaba en el ejemplo
	// c.JSON(200, gin.H{
	// 	"Id":       u.Id,
	// 	"Name":     u.Name,
	// 	"LastName": u.LastName,
	// 	// Email: "gian@example.com",
	// 	// Age: 26,
	// 	// Height: 1.73,
	// 	// Activo: true,∫
	// 	// Date: "03-03-2022"
	// })
}

func main() {
	//var users []User
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gianfranco",
		})
	})

	jsonToStruct()

	//router.GET("/users", GetAll)

	router.Run()
}
