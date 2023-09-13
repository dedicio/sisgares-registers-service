package usecase

import (
	"github.com/dedicio/sisgares-registers-service/internal/registers/dto"
	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type FindProductByIdUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewFindProductByIdUseCase(productRepository entity.ProductRepository) *FindProductByIdUseCase {
	return &FindProductByIdUseCase{
		ProductRepository: productRepository,
	}
}

func (uc FindProductByIdUseCase) Execute(id string) (*dto.ProductResponseDto, error) {
	product, err := uc.ProductRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return &dto.ProductResponseDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CategoryId:  product.CategoryId,
		Tags:        product.Tags,
	}, nil
}
