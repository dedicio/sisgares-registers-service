package entity

import (
	"github.com/dedicio/sisgares-registers-service/pkg/utils"
)

type PositionRepository interface {
	FindById(id string) (*Position, error)
	FindAll(companyID string) ([]*Position, error)
	Create(position *Position) error
	Update(position *Position) error
	Delete(id string) error
}

type GroupRepository interface {
	FindById(id string) (*Group, error)
	FindAll(companyID string) ([]*Group, error)
	Create(category *Group) error
	Update(category *Group) error
	Delete(id string) error
}
type Group struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}
type Position struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupId     string `json:"group_id"`
	CompanyId   string `json:"company_id"`
}

func NewPosition(
	name string,
	description string,
	groupId string,
	companyId string,
) *Position {
	return &Position{
		ID:          utils.NewUUID(),
		Name:        name,
		Description: description,
		GroupId:     groupId,
		CompanyId:   companyId,
	}
}

func NewGroup(
	name string,
	companyId string,
) *Group {
	return &Group{
		ID:        utils.NewUUID(),
		Name:      name,
		CompanyId: companyId,
	}
}
