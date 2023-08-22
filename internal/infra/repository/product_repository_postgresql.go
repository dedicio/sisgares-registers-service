package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/domain/entity"
)

type ProductRepositoryPostgresql struct {
	DB *sql.DB
}

func NewProductRepositoryPostgresql(db *sql.DB) *ProductRepositoryPostgresql {
	return &ProductRepositoryPostgresql{
		DB: db,
	}
}

func (pr *ProductRepositoryPostgresql) Create(product *entity.Product) error {
	_, err := pr.DB.Exec(
		`INSERT INTO products (id, name, description, price, image, category_id, tags, company_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Image,
		product.CategoryId,
		product.Tags,
		product.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryPostgresql) FindAll() ([]*entity.Product, error) {
	rows, err := pr.DB.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image,
			&product.CategoryId,
			&product.Tags,
			&product.CompanyId,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
