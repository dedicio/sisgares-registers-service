package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type FindCategoryByIdUseCase struct {
	Repository entity.CategoryRepository
	CompanyID  string
}

func NewFindCategoryByIdUseCase(
	categoryRepository entity.CategoryRepository,
	companyID string,
) *FindCategoryByIdUseCase {
	return &FindCategoryByIdUseCase{
		Repository: categoryRepository,
		CompanyID:  companyID,
	}
}

func (uc *FindCategoryByIdUseCase) Execute(id string) (*dto.CategoryResponseDto, error) {
	category, err := uc.Repository.FindById(uc.CompanyID, id)
	if err != nil {
		return nil, err
	}

	output := &dto.CategoryResponseDto{
		ID:   category.ID,
		Name: category.Name,
	}

	return output, nil
}
