package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/users"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/pkg/web"
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

type UserRequest struct {
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

type UserRequestPatch struct {
	Name string `json:"name" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(service users.Service) *User {
	return &User{
		service: service,
	}
}

// Store User godoc
// @Summary  Store user
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    user	  body      UserRequest		true  "User to Store"
// @Success  200      {object}  web.Response	"New user"
// @Failure  401      {object}  web.Response	"Unauthorized"
// @Failure  400      {object}  web.Response	"BadRequest"
// @Router   /users   [POST]
func (c *User) Store(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene permisos para realizar la petición solicitada"))
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
		return
	}

	u, err := c.service.Store(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(200, u, ""))
}

// Listusers godoc
// @Summary  Showlist users
// @Tags Users
// @Produce json
// @Param token header string  		  true "token"
// @Success 200 {object} web.Response "List users"
// @Failure 401 {object} web.Response "Unauthorized"
// @Failure 404 {object} web.Response "Not found users"
// @Failure 500 {object} web.Response "Internal server error"
// @Router /users [GET]
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

// Update user godoc
// @Summary  Update user
// @Tags     Users
// @Accept   json
// @Produce   json
// @Param    id       path      int             true   "Id user"
// @Param    token    header    string          true  "token"
// @Param    user	  body 	    UserRequest		true  "User to Update"
// @Success  200      {object}  web.Response 	"Update user"
// @Failure  401      {object}  web.Response 	"Unauthorized"
// @Failure  400	  {object}  web.Response    "BadRequest"
// @Failure  404      {object}  web.Response 	"Not found user"
// @Router   /users/{id}   [PUT]
func (c *User) Update(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No tiene permisos para realizar la petición solicitada"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "ID invalido"))
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Nombre de usuario no puede ser nulo"))
		return
	}

	if req.LastName == "" {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Apellido de usuario no puede ser nulo"))
		return
	}

	// !- Y asi para cada campo xd

	u, err := c.service.Update(int(id), req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.Date)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(200, u, ""))
}

// Delete user
// @Summary Delete user
// @Tags     Users
// @Param    id       path      int             true   "user id"
// @Param    token    header    string          true  "token"
// @Success  200
// @Failure  401      {object}  web.Response 	"Unauthorized"
// @Failure  400      {object}  web.Response 	"BadRequest"
// @Failure  404      {object}  web.Response 	"Not found"
// @Router   /users/{id}   [DELETE]
func (c *User) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "No tiene permisos para realizar la petición solicitada"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID invalido"))
		return
	}

	err = c.service.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(200, "Usuario eliminado correctamente", ""))
}

// UpdateName godoc
// @Summary Update Name
// @Tags     Users
// @Accept   json
// @Produce   json
// @Param    token    header    string           true  "token"
// @Param    id       path      int              true   "user id"
// @Param    name     body      UserRequestPatch true   "user name"
// @Success  200      {object}  web.Response 	 "User"
// @Failure  401      {object}  web.Response 	 "Unauthorized"
// @Failure  400      {object}  web.Response     "BadRequest"
// @Failure  404      {object}  web.Response 	 "Not found"
// @Router   /users/{id} [PATCH]
func (c *User) UpdateName(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "No tiene permisos para realizar la petición solicitada"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID invalido"))
		return
	}

	var req UserRequestPatch
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
		return
	}

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Nombre de usuario no puede ser nulo"))
		return
	}

	u, err := c.service.UpdateName(int(id), req.Name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(200, u, ""))
	return
}
