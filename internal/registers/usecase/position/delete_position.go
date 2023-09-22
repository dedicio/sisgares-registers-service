package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type DeletePositionUseCase struct {
	Repository entity.PositionRepository
}

func NewDeletePositionUseCase(positionRepository entity.PositionRepository) *DeletePositionUseCase {
	return &DeletePositionUseCase{
		Repository: positionRepository,
	}
}

func (uc DeletePositionUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
