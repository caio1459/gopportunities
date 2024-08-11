package router

import (
	"github.com/caio1459/gopportunities/src/controllers"
	"github.com/gin-gonic/gin"
)

const (
	basePath string = "/api/v1"
)

// initializeRoutes Inicia as rotas
func initializeRoutes(router *gin.Engine) {
	//Iniciando os controllers
	controllers.Init()
	//Defini um agrupamento de rotas
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", controllers.ListOpenings)
		v1.GET("/opening", controllers.ShowOpening)
		v1.POST("/openings", controllers.CreateOpening)
		v1.DELETE("/openings", controllers.DeleteOpening)
	}
}
