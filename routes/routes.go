package routes

import (
	"github.com/cristopherecruz/go-desenvolvendo-api-rest/controllers"
	"github.com/cristopherecruz/go-desenvolvendo-api-rest/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequest() {

	router := gin.Default()

	router.Use(cors.Default())

	router.Use(middleware.HeaderMiddleware)

	router.GET("/", controllers.Home)
	router.GET("/api/personalidades", controllers.BuscarTodasPersonalidades)
	router.GET("/api/personalidades/:id", controllers.BuscarPersonalidadePorId)
	router.POST("/api/personalidades", controllers.CriarPersonalidade)
	router.DELETE("/api/personalidades/:id", controllers.DeletarPersonalidade)
	router.PUT("/api/personalidades/:id", controllers.EditarPersonalidade)

	log.Fatal(router.Run())
}
