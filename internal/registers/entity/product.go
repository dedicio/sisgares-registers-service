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
}

type CategoryRepository interface {
	FindById(id string) (*Category, error)
	FindAll() ([]*Category, error)
	Create(category *Category) error
	Update(category *Category) error
	Delete(id string) error
	FindProductsByCategoryId(categoryId string) ([]*Product, error)
}

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

func NewCategory(name string, companyId string) *Category {
	id := utils.NewUUID()
	return &Category{
		ID:        id,
		Name:      name,
		CompanyId: companyId,
	}
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
