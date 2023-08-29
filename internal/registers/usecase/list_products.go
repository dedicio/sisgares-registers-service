package usecase

import "github.com/dedicio/sisgares-registers-service/internal/registers/entity"

type ListProductsOutputDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	CategoryId  string   `json:"category_id"`
	Tags        []string `json:"tags"`
	CompanyId   string   `json:"company_id"`
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
