package usecase

import "github.com/dedicio/sisgares-registers-service/internal/registers/entity"

type UpdateProductInputDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	CategoryId  string   `json:"category_id"`
	Tags        []string `json:"tags"`
	CompanyId   string   `json:"company_id"`
}

type UpdateProductUseCase struct {
	Repository entity.ProductRepository
}

func NewUpdateProductUseCase(productRepository entity.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		Repository: productRepository,
	}
}

func (uc *UpdateProductUseCase) Execute(input UpdateProductInputDto) error {
	product, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Image = input.Image
	product.CategoryId = input.CategoryId
	product.Tags = input.Tags
	product.CompanyId = input.CompanyId

	err = uc.Repository.Update(product)
	if err != nil {
		return err
	}

	return nil
}
