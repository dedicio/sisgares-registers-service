package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type DeleteCategoryUseCase struct {
	Repository entity.CategoryRepository
}

func NewDeleteCategoryUseCase(categoryRepository entity.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		Repository: categoryRepository,
	}
}

func (uc *DeleteCategoryUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
