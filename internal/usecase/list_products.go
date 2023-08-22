package usecase

import "github.com/dedicio/sisgares-registers-service/internal/domain/entity"

type ListProductsOutputDto struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Image       string
	CategoryId  string
	Tags        []string
	CompanyId   string
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (uc ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := uc.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*ListProductsOutputDto
	for _, product := range products {
		output = append(output, &ListProductsOutputDto{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
			CategoryId:  product.CategoryId,
			Tags:        product.Tags,
			CompanyId:   product.CompanyId,
		})
	}

	return output, nil
}
