package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (pr *ProductRepositoryMysql) FindById(id string) (*entity.Product, error) {
	var product entity.Product

	sql := `SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL`
	err := pr.DB.QueryRow(sql, id).Scan(
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

	return &product, nil
}

func (pr *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	sql := `SELECT * FROM products WHERE deleted_at IS NULL`
	rows, err := pr.DB.Query(sql)
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

func (pr *ProductRepositoryMysql) Create(product *entity.Product) error {
	sql := `INSERT INTO products (id, name, description, price, image, category_id, tags, company_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := pr.DB.Exec(
		sql,
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

func (pr *ProductRepositoryMysql) Update(product *entity.Product) error {
	sql := `UPDATE products SET name = $1, description = $2, price = $3, image = $4, category_id = $5, tags = $6, company_id = $7 WHERE id = $8`
	_, err := pr.DB.Exec(
		sql,
		product.Name,
		product.Description,
		product.Price,
		product.Image,
		product.CategoryId,
		product.Tags,
		product.CompanyId,
		product.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryMysql) Delete(id string) error {
	sql := `UPDATE products SET deleted_at = NOW() WHERE id = $1`
	_, err := pr.DB.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryMysql) FindByCategoryId(categoryId string) ([]*entity.Product, error) {
	sql := `SELECT * FROM products WHERE category_id = $1 AND deleted_at IS NULL`
	rows, err := pr.DB.Query(sql, categoryId)
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
