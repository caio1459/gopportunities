package models

import (
	"fmt"
	"time"

	"github.com/caio1459/gopportunities/src/helpers"
	"gorm.io/gorm"
)

// Espelho da tabela
type Opening struct {
	Role       string
	Company    string
	Location   string
	Remote     bool
	Link       string
	Salary     float32
	gorm.Model //Integar as propriedades do gorm no struct
}

// Retorno em json
type OpeningResponse struct {
	Id        uint      `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    float32   `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"` //Obriga a pessoa a escolher um valor já que o valor zero é falso
	Link     string  `json:"link"`
	Salary   float32 `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" && c.Company == "" && c.Location == "" && c.Link == "" && c.Remote == nil && c.Salary <= 0 {
		return fmt.Errorf("corpo da requisição está vazio")
	}
	if c.Role == "" {
		return helpers.ErrParamsIsRequired("role", "string")
	}
	if c.Company == "" {
		return helpers.ErrParamsIsRequired("company", "string")
	}
	if c.Location == "" {
		return helpers.ErrParamsIsRequired("location", "string")
	}
	if c.Remote == nil {
		return helpers.ErrParamsIsRequired("remote", "boolean")
	}
	if c.Link == "" {
		return helpers.ErrParamsIsRequired("link", "string")
	}
	if c.Salary <= 0 {
		return helpers.ErrParamsIsRequired("salary", "float")
	}
	return nil
}
