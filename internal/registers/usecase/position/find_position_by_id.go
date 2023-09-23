package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type FindPositionByIdUseCase struct {
	Repository entity.PositionRepository
}

func NewFindPositionByIdUseCase(positionRepository entity.PositionRepository) *FindPositionByIdUseCase {
	return &FindPositionByIdUseCase{
		Repository: positionRepository,
	}
}

func (uc FindPositionByIdUseCase) Execute(id string) (*dto.PositionResponseDto, error) {
	position, err := uc.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	output := &dto.PositionResponseDto{
		ID:          position.ID,
		Name:        position.Name,
		Description: position.Description,
		GroupId:     position.GroupId,
	}

	return output, nil
}
