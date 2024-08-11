package controllers

import (
	"github.com/caio1459/gopportunities/src/config"
	"gorm.io/gorm"
)

var (
	logger        *config.Logger
	db            *gorm.DB
	err           error
	total         int64
	returnedTotal int
)

func Init() {
	logger = config.NewLogger("controller")
	db = config.GetMySQL()
}
