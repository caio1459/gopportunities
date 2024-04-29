package router

import (
	"github.com/caio1459/gopportunities/src/controllers"
	"github.com/gin-gonic/gin"
)

// Inicia as rotas
func initializeRoutes(router *gin.Engine) {
	//Iniciando os controllers
	controllers.Init()
	//Defini um agrupamento de rotas
	openingRoute := router.Group("/api/v1/openings")
	{
		openingRoute.GET("/", controllers.ShowOpening)
		openingRoute.POST("/", controllers.CreateOpening)
		openingRoute.PUT("/", controllers.UpdateOpening)
		openingRoute.DELETE("/", controllers.DeleteOpening)
	}
}
