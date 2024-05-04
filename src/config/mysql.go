package config

import (
	"fmt"
	"os"

	"github.com/caio1459/gopportunities/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	// Cria o dns de conexão com banco
	dns := fmt.Sprintf("%v:@/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_NAME"))

	//Criando conexão com banco de dados utilizando o gorm
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		logger.ErrF("Erro ao criar conexão com o MySql: %v", err)
		return nil, err
	}

	//Cria as migrations do banco de dados
	if err = db.AutoMigrate(&models.Opening{}); err != nil {
		logger.ErrF("Erro ao criar migrations: %v", err)
		return nil, err
	}

	return db, nil
}

func GetMySQL() *gorm.DB {
	return db
}
