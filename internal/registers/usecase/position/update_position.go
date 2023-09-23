package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type UpdatePositionUseCase struct {
	Repository entity.PositionRepository
}

func NewUpdatePositionUseCase(positionRepository entity.PositionRepository) *UpdatePositionUseCase {
	return &UpdatePositionUseCase{
		Repository: positionRepository,
	}
}

func (uc UpdatePositionUseCase) Execute(input dto.PositionDto) error {
	position, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	position.Name = input.Name
	position.Description = input.Description
	position.GroupId = input.GroupId
	position.CompanyId = input.CompanyId

	err = uc.Repository.Update(position)
	if err != nil {
		return err
	}

	return nil
}
