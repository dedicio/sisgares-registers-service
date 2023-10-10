package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListCategoriesUseCase struct {
	Repository entity.CategoryRepository
	CompanyID  string
}

func NewListCategoriesUseCase(
	categoryRepository entity.CategoryRepository,
	companyID string,
) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{
		Repository: categoryRepository,
		CompanyID:  companyID,
	}
}

func (uc ListCategoriesUseCase) Execute() ([]*dto.CategoryResponseDto, error) {
	categories, err := uc.Repository.FindAll(uc.CompanyID)
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
