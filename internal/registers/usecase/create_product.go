package usecase

import "github.com/dedicio/sisgares-registers-service/internal/registers/entity"

type CreateProductInputDto struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	CategoryId  string   `json:"category_id"`
	Tags        []string `json:"tags"`
	CompanyId   string   `json:"company_id"`
}

type CreateProductOutputDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	CategoryId  string   `json:"category_id"`
	Tags        []string `json:"tags"`
}

type CreateProductUseCase struct {
	Repository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		Repository: productRepository,
	}
}

func (uc *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
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

	output := &CreateProductOutputDto{
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
