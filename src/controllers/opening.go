package controllers

import (
	"net/http"

	"github.com/caio1459/gopportunities/src/helpers"
	"github.com/caio1459/gopportunities/src/models"
	s "github.com/caio1459/gopportunities/src/services"
	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	request := models.CreateOpeningRequest{}
	//Pega o json da requisição e passa para o struct
	if err = ctx.BindJSON(&request); err != nil {
		logger.ErrF("Erro de conversão: %v", err.Error())
		s.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//Valida se os campos estão corretos
	if err = request.Validate(); err != nil {
		logger.ErrF("Erro de validação: %v", err.Error())
		s.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := models.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err = db.Create(&opening).Error; err != nil {
		logger.ErrF("Erro ao criar vaga: %v", err.Error())
		s.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	s.SendSuccess(ctx, "create-opening", opening)
}

func ShowOpenings(ctx *gin.Context) {
	openings := []models.Opening{}

	if err = db.Find(&openings).Error; err != nil {
		s.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	s.SendSuccess(ctx, "list-openings", openings)
}

func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT",
	})
}

func DeleteOpening(ctx *gin.Context) {
	//Pega o parametro da url
	id := ctx.Query("id")
	if id == "" {
		s.SendError(ctx, http.StatusBadRequest, helpers.ErrParamsIsRequired("id", "queryParameter").Error())
		return
	}

	opening := models.Opening{}
	//Verifca se existe um registro no banco
	if err = db.First(&opening, id).Error; err != nil {
		s.SendError(ctx, http.StatusNotFound, "Vaga inexistente")
		return
	}

	if err = db.Delete(&opening).Error; err != nil {
		s.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	s.SendSuccess(ctx, "delete-opening", opening)
}
