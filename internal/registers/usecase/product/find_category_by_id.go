package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type FindCategoryByIdUseCase struct {
	Repository entity.CategoryRepository
}

func NewFindCategoryByIdUseCase(categoryRepository entity.CategoryRepository) *FindCategoryByIdUseCase {
	return &FindCategoryByIdUseCase{
		Repository: categoryRepository,
	}
}

func (uc *FindCategoryByIdUseCase) Execute(id string) (*dto.CategoryResponseDto, error) {
	category, err := uc.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	output := &dto.CategoryResponseDto{
		ID:   category.ID,
		Name: category.Name,
	}

	return output, nil
}
