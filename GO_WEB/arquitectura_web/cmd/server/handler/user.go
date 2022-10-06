package handler

import (
	"net/http"
	"os"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/internal/users"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web/pkg/web"
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
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene permisos para realizar la petición solicitada"))
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(500, nil, err.Error()))
		return
	}

	u, err := c.service.Create(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error()))
	}
	ctx.JSON(http.StatusOK, web.NewResponse(200, u, ""))
}

func (c *User) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene permisos para realizar la petición solicitada"))
		return
	}

	u, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error()))
		return
	}

	if len(u) == 0 {
		ctx.JSON(http.StatusOK, web.NewResponse(404, nil, "No hay usuarios almacenados"))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(200, u, ""))
	return
}
