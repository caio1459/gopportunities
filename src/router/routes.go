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
	openingRoute := router.Group("/api/v1")
	{
		openingRoute.GET("/openings", controllers.ShowOpenings)
		openingRoute.POST("/openings", controllers.CreateOpening)
		openingRoute.PUT("/openings", controllers.UpdateOpening)
		openingRoute.DELETE("/openings", controllers.DeleteOpening)
	}
}
