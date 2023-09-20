package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListProductsByCategoryUseCase struct {
	CategoryRepository entity.CategoryRepository
}

func NewListProductsByCategoryUseCase(productRepository entity.CategoryRepository) *ListProductsByCategoryUseCase {
	return &ListProductsByCategoryUseCase{
		CategoryRepository: productRepository,
	}
}

func (uc ListProductsByCategoryUseCase) Execute(categoryId string) ([]*dto.ProductResponseDto, error) {
	products, err := uc.CategoryRepository.FindProductsByCategoryId(categoryId)
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
