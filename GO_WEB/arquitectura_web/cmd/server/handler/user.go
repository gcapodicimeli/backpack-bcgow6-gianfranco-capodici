package handler

import (
	"net/http"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(service users.Service) *User {
	return &User{
		service: service,
	}
}

func (c *User) Create(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "12345" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := c.service.Create(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, u)
}

func (c *User) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "12345" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	u, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, u)
	return
}
