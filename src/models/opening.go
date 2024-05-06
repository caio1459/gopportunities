package models

import (
	"fmt"

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
	Id       uint    `json:"id"`
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   bool    `json:"remote"`
	Link     string  `json:"link"`
	Salary   float32 `json:"salary"`
}

type OpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"` //Obriga a pessoa a escolher um valor já que o valor zero é falso
	Link     string  `json:"link"`
	Salary   float32 `json:"salary"`
}

func (c *OpeningRequest) Validate() error {
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

// Seta os valores da requeste na opening de acordo com operação de insert ou update
func (r OpeningRequest) SetOpening(stage string, opening *Opening) (*Opening, error) {
	if stage == "insert" {
		return &Opening{
			Role:     r.Role,
			Company:  r.Company,
			Location: r.Location,
			Remote:   *r.Remote,
			Link:     r.Link,
			Salary:   r.Salary,
		}, nil
	}
	if stage == "update" && opening != nil {
		if r.Role != "" {
			opening.Role = r.Role
		}
		if r.Company != "" {
			opening.Company = r.Company
		}
		if r.Location != "" {
			opening.Location = r.Location
		}
		if r.Remote != nil {
			opening.Remote = *r.Remote
		}
		if r.Link != "" {
			opening.Link = r.Link
		}
		if r.Salary > 0 {
			opening.Salary = r.Salary
		}
		return opening, nil
	}
	return nil, helpers.ErrParamsIsRequired("update or insert", "string")
}

// Método para serializar uma abertura para uma resposta JSON
func (o Opening) SetOpeningResponse() OpeningResponse {
	return OpeningResponse{
		Id:       o.ID,
		Role:     o.Role,
		Company:  o.Company,
		Location: o.Location,
		Remote:   o.Remote,
		Link:     o.Link,
		Salary:   o.Salary,
	}
}

// Função para serializar uma abertura para uma resposta em um array de JSON
func SetOpeningsResponse(openings []Opening) []OpeningResponse {
	responses := []OpeningResponse{}
	for _, opening := range openings {
		response := opening.SetOpeningResponse()
		responses = append(responses, response)
	}
	return responses
}
