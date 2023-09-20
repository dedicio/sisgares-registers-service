package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListCategoriesUseCase struct {
	Repository entity.CategoryRepository
}

func NewListCategoriesUseCase(categoryRepository entity.CategoryRepository) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{
		Repository: categoryRepository,
	}
}

func (uc ListCategoriesUseCase) Execute() ([]*dto.CategoryResponseDto, error) {
	categories, err := uc.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*dto.CategoryResponseDto
	for _, category := range categories {
		output = append(output, &dto.CategoryResponseDto{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return output, nil
}
