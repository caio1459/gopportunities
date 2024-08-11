package controllers

import (
	"net/http"

	"github.com/caio1459/gopportunities/src/helpers"
	"github.com/caio1459/gopportunities/src/models"
	r "github.com/caio1459/gopportunities/src/responses"
	s "github.com/caio1459/gopportunities/src/services"
	"github.com/gin-gonic/gin"
)

// Cria uma Vaga
func CreateOpening(ctx *gin.Context) {
	req := models.OpeningRequest{}
	//Pega o json da requisição e passa para o struct
	if err = ctx.BindJSON(&req); err != nil {
		logger.ErrF("Erro de conversão: %v", err.Error())
		r.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//Valida se os campos estão corretos
	if err = req.Validate(); err != nil {
		logger.ErrF("Erro de validação: %v", err.Error())
		r.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening, err := req.SetOpening("insert", nil)
	if err != nil {
		logger.ErrF("Erro de conversão: %v", err.Error())
		r.SendError(ctx, http.StatusInternalServerError, "Erro interno")
		return
	}

	//Executa o Insert
	if err = db.Create(&opening).Error; err != nil {
		logger.ErrF("Erro ao criar vaga: %v", err.Error())
		r.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := opening.SetOpeningResponse()
	r.SendSuccess(ctx, "create-opening", res)
}
// Lista uma vaga sem paginação
// func ShowOpening(ctx *gin.Context) {
// 	//Pega o parametro da url
// 	id := ctx.Query("id")
// 	if id == "" {
// 		s.SendError(ctx, http.StatusBadRequest, helpers.ErrParamsIsRequired("id", "queryParameter").Error())
// 		return
// 	}

// 	opening := models.Opening{}
// 	if err = db.First(&opening, id).Error; err != nil {
// 		s.SendError(ctx, http.StatusNotFound, "Vaga inexistente")
// 		return
// 	}

// 	res := opening.SetOpeningResponse()
// 	s.SendSuccess(ctx, "show-opening", res)
// }
// Lista todas as vagas
func ListOpenings(ctx *gin.Context) {
	openings := []models.Opening{}
	// Obtém os parâmetros de paginação
	pagination := s.ParsePaginationParams(ctx)
	offset := pagination.CalculateOffset()

	// Conta o número total de registros
	if err = db.Model(&models.Opening{}).Count(&total).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Executa a consulta com limite e offset para paginação
	if err = db.Limit(pagination.PageSize).Offset(offset).Find(&openings).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Calcula o total de registros retornados
	returnedTotal = len(openings)

	res := s.CreatePaginatedResponse(
		models.SetOpeningsResponse(openings),
		pagination.Page,
		pagination.PageSize,
		total,
		returnedTotal,
	)
	r.SendSuccess(ctx, "list-openings", res)
}

// Lista uma vaga
func ShowOpening(ctx *gin.Context) {
	//Pega o parametro da url
	id := ctx.Query("id")
	if id == "" {
		r.SendError(ctx, http.StatusBadRequest, helpers.ErrParamsIsRequired("id", "queryParameter").Error())
		return
	}

	opening := models.Opening{}
	if err = db.First(&opening, id).Error; err != nil {
		r.SendError(ctx, http.StatusNotFound, "Vaga inexistente")
		return
	}

	res := opening.SetOpeningResponse()
	r.SendSuccess(ctx, "show-opening", res)
}

// Atualiza uma vaga
func UpdateOpening(ctx *gin.Context) {
	req := models.OpeningRequest{}
	if err = ctx.BindJSON(&req); err != nil {
		r.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err = req.Validate(); err != nil {
		logger.ErrF("Erro de validação: %v", err.Error())
		r.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		r.SendError(ctx, http.StatusBadRequest, helpers.ErrParamsIsRequired("id", "queryParameter").Error())
		return
	}

	//Verifica se existe a vaga
	opening := models.Opening{}
	if err = db.First(&opening, id).Error; err != nil {
		r.SendError(ctx, http.StatusNotFound, "Vaga inexistente")
		return
	}

	//Cria uma nova vaga de acordo com a requisição
	newOpening, err := req.SetOpening("update", &opening)
	if err != nil {
		logger.ErrF("Erro de conversão: %v", err.Error())
		r.SendError(ctx, http.StatusInternalServerError, "Erro interno")
		return
	}

	//Realiza o update
	if err = db.Save(&newOpening).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := newOpening.SetOpeningResponse()
	r.SendSuccess(ctx, "update-opening", res)
}

// Deleta uma vaga
func DeleteOpening(ctx *gin.Context) {
	//Pega o parametro da url
	id := ctx.Query("id")
	if id == "" {
		r.SendError(ctx, http.StatusBadRequest, helpers.ErrParamsIsRequired("id", "queryParameter").Error())
		return
	}

	opening := models.Opening{}
	//Verifca se existe um registro no banco
	if err = db.First(&opening, id).Error; err != nil {
		r.SendError(ctx, http.StatusNotFound, "Vaga inexistente")
		return
	}

	if err = db.Delete(&opening).Error; err != nil {
		r.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	res := opening.SetOpeningResponse()
	r.SendSuccess(ctx, "delete-opening", res)
}
