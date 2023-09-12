package entity

import (
	"github.com/dedicio/sisgares-registers-service/pkg/utils"
)

type ProductRepository interface {
	FindById(id string) (*Product, error)
	FindAll() ([]*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id string) error
	FindByCategoryId(categoryId string) ([]*Product, error)
}

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryId  string  `json:"category_id"`
	CompanyId   string  `json:"company_id"`
}

func NewProduct(
	name string,
	description string,
	price float64,
	image string,
	categoryId string,
	companyId string,
) *Product {
	id := utils.NewUUID()
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Image:       image,
		CategoryId:  categoryId,
		CompanyId:   companyId,
	}
}
