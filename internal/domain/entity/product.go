package entity

import (
	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
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
