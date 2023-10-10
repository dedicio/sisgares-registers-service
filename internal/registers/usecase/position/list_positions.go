package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListPositionsUseCase struct {
	Repository entity.PositionRepository
}

func NewListPositionsUseCase(positionRepository entity.PositionRepository) *ListPositionsUseCase {
	return &ListPositionsUseCase{
		Repository: positionRepository,
	}
}

func (uc ListPositionsUseCase) Execute(companyID string) ([]*dto.PositionResponseDto, error) {
	positions, err := uc.Repository.FindAll(companyID)
	if err != nil {
		return nil, err
	}

	var output []*dto.PositionResponseDto
	for _, position := range positions {
		output = append(output, &dto.PositionResponseDto{
			ID:          position.ID,
			Name:        position.Name,
			Description: position.Description,
			GroupId:     position.GroupId,
		})
	}

	return output, nil
}
