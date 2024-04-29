package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() error {
	logger := GetLogger("config")

	//Iniciando variaveis de ambiemte
	if err = godotenv.Load(); err != nil {
		logger.ErrF("Erro ao carregar variaveis de ambiente: %v", err)
		return err
	}

	//Inicando banco de dados
	db, err = InitializeMySQL()
	if err != nil {
		logger.ErrF("Erro ao iniciar o banco de dados: %v", err)
		return err
	}
	return nil
}
