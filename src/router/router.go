package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Exemplo de API com Gin e Swagger
// @version 1.0
// @description Esta é uma API de exemplo.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
func Initialize() {
	//Inicia o router
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Inicia as rotas
	initializeRoutes(router)
	//Inicia o server
	router.Run("localhost:3400")
}
