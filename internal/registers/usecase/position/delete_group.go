package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type DeleteGroupUseCase struct {
	Repository entity.GroupRepository
}

func NewDeleteGroupUseCase(groupRepository entity.GroupRepository) *DeleteGroupUseCase {
	return &DeleteGroupUseCase{
		Repository: groupRepository,
	}
}

func (uc DeleteGroupUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
