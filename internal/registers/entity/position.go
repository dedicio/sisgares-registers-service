package entity

import "github.com/google/uuid"

type Type struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}
type Position struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TypeId      string `json:"type_id"`
	CompanyId   string `json:"company_id"`
}

func NewPosition(
	name string,
	description string,
	typeId string,
	companyId string,
) *Position {
	return &Position{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		TypeId:      typeId,
		CompanyId:   companyId,
	}
}
