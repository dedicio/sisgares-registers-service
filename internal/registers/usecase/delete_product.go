package usecase

import "github.com/dedicio/sisgares-registers-service/internal/registers/entity"

type DeleteProductUseCase struct {
	Repository entity.ProductRepository
}

func NewDeleteProductUseCase(productRepository entity.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		Repository: productRepository,
	}
}

func (uc *DeleteProductUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
