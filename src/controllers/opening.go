package controllers

import (
	"net/http"

	"github.com/caio1459/gopportunities/src/models"
	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	request := models.CreateOpeningRequest{}
	//Pega o json da requisição e passa para o struct
	if err = ctx.BindJSON(&request); err != nil {
		logger.ErrF("Erro de conversão: %v", err.Error())
		return
	}
	//Valida se os campos estão corretos
	if err = request.Validate(); err != nil {
		logger.ErrF("Erro de validação: %v", err.Error())
		return
	}

	if err = db.Create(&request).Error; err != nil {
		logger.ErrF("Erro ao criar vaga: %v", err.Error())
		return
	}
}

func ShowOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET",
	})
}

func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT",
	})
}

func DeleteOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE",
	})
}
