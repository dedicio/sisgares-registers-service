package usecase

import "github.com/dedicio/sisgares-registers-service/internal/domain/entity"

type CreateProductInputDto struct {
	Name        string
	Description string
	Price       float64
	Image       string
	CategoryId  string
	Tags        []string
	CompanyId   string
}

type CreateProductOutputDto struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Image       string
	CategoryId  string
	Tags        []string
	CompanyId   string
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (uc CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(
		input.Name,
		input.Description,
		input.Price,
		input.Image,
		input.CategoryId,
		input.Tags,
		input.CompanyId,
	)

	err := uc.ProductRepository.Create(product)
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
		CompanyId:   product.CompanyId,
	}

	return output, nil
}
