package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type UpdateGroupUseCase struct {
	Repository entity.GroupRepository
}

func NewUpdateGroupUseCase(groupRepository entity.GroupRepository) *UpdateGroupUseCase {
	return &UpdateGroupUseCase{
		Repository: groupRepository,
	}
}

func (uc UpdateGroupUseCase) Execute(input dto.GroupDto) error {
	group, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	group.Name = input.Name
	group.CompanyId = input.CompanyId

	err = uc.Repository.Update(group)
	if err != nil {
		return err
	}

	return nil
}
