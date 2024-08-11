package services

import (
	"strconv"

	r "github.com/caio1459/gopportunities/src/responses"
	"github.com/gin-gonic/gin"
)

type PaginationParams struct {
	Page     int
	PageSize int
}

func ParsePaginationParams(ctx *gin.Context) PaginationParams {
	pageQuery := ctx.DefaultQuery("page", "1")
	pageSizeQuery := ctx.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 1 {
		page = 1 // Página padrão
	}

	pageSize, err := strconv.Atoi(pageSizeQuery)
	if err != nil || pageSize < 1 {
		pageSize = 10 // Tamanho da página padrão
	}

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

// CalculateOffset calcula o offset com base na página e no tamanho da página.
func (p PaginationParams) CalculateOffset() int {
	return (p.Page - 1) * p.PageSize
}

// CalculateTotalPages calcula o número total de páginas com base no total de registros e no tamanho da página.
func CalculateTotalPages(total int64, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return int((total + int64(pageSize) - 1) / int64(pageSize)) // Arredonda para cima
}

// CreatePaginatedResponse cria uma resposta paginada usando o struct padrão.
func CreatePaginatedResponse[T any](data T, page, pageSize int, total int64, returnedTotal int) r.PaginatedResponse[T] {
	return r.PaginatedResponse[T]{
		Data:          data,
		Page:          page,
		PageSize:      pageSize,
		Total:         total,
		TotalPages:    CalculateTotalPages(total, pageSize),
		ReturnedTotal: returnedTotal,
	}
}
