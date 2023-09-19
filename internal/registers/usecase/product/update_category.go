package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type UpdateCategoryUseCase struct {
	Repository entity.CategoryRepository
}

func NewUpdateCategoryUseCase(categoryRepository entity.CategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		Repository: categoryRepository,
	}
}

func (uc *UpdateCategoryUseCase) Execute(input dto.CategoryDto) error {
	category, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	category.Name = input.Name
	category.CompanyId = input.CompanyId

	err = uc.Repository.Update(category)
	if err != nil {
		return err
	}

	return nil
}
