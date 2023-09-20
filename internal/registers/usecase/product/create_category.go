package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type CreateCategoryUseCase struct {
	Repository entity.CategoryRepository
}

func NewCreateCategoryUseCase(categoryRepository entity.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		Repository: categoryRepository,
	}
}

func (uc *CreateCategoryUseCase) Execute(input dto.CategoryDto) (*dto.CategoryDto, error) {
	category := entity.NewCategory(input.Name, input.CompanyId)

	err := uc.Repository.Create(category)
	if err != nil {
		return nil, err
	}

	output := &dto.CategoryDto{
		ID:        category.ID,
		Name:      category.Name,
		CompanyId: category.CompanyId,
	}

	return output, nil
}
