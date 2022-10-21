package main

import (
	"os"

	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/cmd/server/handler"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/docs"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/internal/users"
	"github.com/gcapodicimeli/backpack-bcgow6-gianfranco-capodici/arquitectura_web_testing/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Arquitectura del bootcamp - API
// @version         1.0
// @description     This is a simple API development by Gianfranco Capodici.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   Gianfranco Capodici
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ur := router.Group("/users")
	ur.POST("/", u.Store)
	ur.GET("/", u.GetAll)
	ur.PUT("/:id", u.Update)
	ur.DELETE("/:id", u.Delete)
	ur.PATCH("/:id", u.UpdateName)

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
