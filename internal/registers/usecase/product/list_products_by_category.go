package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ListProductsByCategoryUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsByCategoryUseCase(productRepository entity.ProductRepository) *ListProductsByCategoryUseCase {
	return &ListProductsByCategoryUseCase{
		ProductRepository: productRepository,
	}
}

func (uc ListProductsByCategoryUseCase) Execute(categoryId string) ([]*dto.ProductResponseDto, error) {
	products, err := uc.ProductRepository.FindByCategoryId(categoryId)
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
			Tags:        product.Tags,
		})
	}

	return output, nil
}
