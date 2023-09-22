package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListGroupsUseCase struct {
	Repository entity.GroupRepository
}

func NewListGroupsUseCase(groupRepository entity.GroupRepository) *ListGroupsUseCase {
	return &ListGroupsUseCase{
		Repository: groupRepository,
	}
}

func (uc ListGroupsUseCase) Execute() ([]*dto.GroupResponseDto, error) {
	groups, err := uc.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*dto.GroupResponseDto
	for _, group := range groups {
		output = append(output, &dto.GroupResponseDto{
			ID:   group.ID,
			Name: group.Name,
		})
	}

	return output, nil
}
