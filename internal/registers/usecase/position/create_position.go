package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type CreatePositionUseCase struct {
	Repository entity.PositionRepository
}

func NewCreatePositionUseCase(positionRepository entity.PositionRepository) *CreatePositionUseCase {
	return &CreatePositionUseCase{
		Repository: positionRepository,
	}
}

func (uc CreatePositionUseCase) Execute(input dto.PositionDto) (*dto.PositionDto, error) {
	position := entity.NewPosition(
		input.Name,
		input.Description,
		input.GroupId,
		input.CompanyId,
	)

	err := uc.Repository.Create(position)
	if err != nil {
		return nil, err
	}

	output := &dto.PositionDto{
		ID:          position.ID,
		Name:        position.Name,
		Description: position.Description,
		GroupId:     position.GroupId,
	}

	return output, nil
}
