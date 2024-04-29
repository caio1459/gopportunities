package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Inicia o router
	router := gin.Default()
	//Inicia as rotas
	initializeRoutes(router)
	//Inicia o server
	router.Run("localhost:3200")
}
