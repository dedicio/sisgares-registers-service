package entity

import (
	"github.com/google/uuid"
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
type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	CategoryId  string   `json:"category_id"`
	Tags        []string `json:"tags"`
	CompanyId   string   `json:"company_id"`
}

func NewProduct(
	name string,
	description string,
	price float64,
	image string,
	categoryId string,
	tags []string,
	companyId string,
) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Image:       image,
		CategoryId:  categoryId,
		Tags:        tags,
		CompanyId:   companyId,
	}
}
