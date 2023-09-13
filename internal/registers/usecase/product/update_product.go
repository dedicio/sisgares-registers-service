package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type UpdateProductUseCase struct {
	Repository entity.ProductRepository
}

func NewUpdateProductUseCase(productRepository entity.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		Repository: productRepository,
	}
}

func (uc *UpdateProductUseCase) Execute(input dto.ProductDto) error {
	product, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Image = input.Image
	product.CategoryId = input.CategoryId
	product.Tags = input.Tags
	product.CompanyId = input.CompanyId

	err = uc.Repository.Update(product)
	if err != nil {
		return err
	}

	return nil
}
