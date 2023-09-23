package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type CreateGroupUseCase struct {
	Repository entity.GroupRepository
}

func NewCreateGroupUseCase(groupRepository entity.GroupRepository) *CreateGroupUseCase {
	return &CreateGroupUseCase{
		Repository: groupRepository,
	}
}

func (uc CreateGroupUseCase) Execute(input dto.GroupDto) (*dto.GroupDto, error) {
	group := entity.NewGroup(
		input.Name,
		input.CompanyId,
	)

	err := uc.Repository.Create(group)
	if err != nil {
		return nil, err
	}

	output := &dto.GroupDto{
		ID:   group.ID,
		Name: group.Name,
	}

	return output, nil
}
