package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type FindGroupByIdUseCase struct {
	Repository entity.GroupRepository
}

func NewFindGroupByIdUseCase(groupRepository entity.GroupRepository) *FindGroupByIdUseCase {
	return &FindGroupByIdUseCase{
		Repository: groupRepository,
	}
}

func (uc FindGroupByIdUseCase) Execute(id string) (*dto.GroupResponseDto, error) {
	group, err := uc.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	output := &dto.GroupResponseDto{
		ID:   group.ID,
		Name: group.Name,
	}

	return output, nil
}
