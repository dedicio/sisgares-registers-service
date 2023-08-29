package usecase

import "github.com/dedicio/sisgares-registers-service/internal/registers/entity"

type ListProductsByCategoryOutputDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	CategoryId  string   `json:"category_id"`
	Tags        []string `json:"tags"`
	CompanyId   string   `json:"company_id"`
}

type ListProductsByCategoryUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsByCategoryUseCase(productRepository entity.ProductRepository) *ListProductsByCategoryUseCase {
	return &ListProductsByCategoryUseCase{
		ProductRepository: productRepository,
	}
}

func (uc ListProductsByCategoryUseCase) Execute(categoryId string) ([]*ListProductsByCategoryOutputDto, error) {
	products, err := uc.ProductRepository.FindByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}

	var output []*ListProductsByCategoryOutputDto
	for _, product := range products {
		output = append(output, &ListProductsByCategoryOutputDto{
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
