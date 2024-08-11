package responses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PaginatedResponse é o struct padrão para respostas paginadas.
type PaginatedResponse[T any] struct {
	Data          T     `json:"data"`          // Dados da página atual
	Page          int   `json:"page"`          // Página atual
	PageSize      int   `json:"pageSize"`      // Tamanho da página
	Total         int64 `json:"total"`         // Total de registros
	TotalPages    int   `json:"totalPages"`    // Total de páginas
	ReturnedTotal int   `json:"returnedTotal"` // Total de registros retornados na página atual
}

func SendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
	})
}

func SendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Sucesso na operação: %v", operation),
		"items":   data,
	})
}
