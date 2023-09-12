package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type CreateProductUseCase struct {
	Repository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		Repository: productRepository,
	}
}

func (uc *CreateProductUseCase) Execute(input dto.ProductDto) (*dto.ProductDto, error) {
	product := entity.NewProduct(
		input.Name,
		input.Description,
		input.Price,
		input.Image,
		input.CategoryId,
		input.Tags,
		input.CompanyId,
	)

	err := uc.Repository.Create(product)
	if err != nil {
		return nil, err
	}

	output := &dto.ProductDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CategoryId:  product.CategoryId,
		Tags:        product.Tags,
	}

	return output, nil
}
