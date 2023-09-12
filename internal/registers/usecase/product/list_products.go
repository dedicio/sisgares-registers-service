package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (uc ListProductsUseCase) Execute() ([]*dto.ProductResponseDto, error) {
	products, err := uc.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*dto.ProductResponseDto
	for _, product := range products {
		output = append(output, &dto.ProductResponseDto{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
			CategoryId:  product.CategoryId,
		})
	}

	return output, nil
}
