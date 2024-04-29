package main

import (
	"github.com/caio1459/gopportunities/src/config"
	r "github.com/caio1459/gopportunities/src/router"
)

func main() {
	logger := config.GetLogger("main")
	//Iniciando as configurações
	err := config.Init()
	if err != nil {
		logger.ErrF("Erro ao iniciar as configurações: %v", err)
	}
	//Iniciando as rotas
	r.Initialize()
}
