package repository

import (
	"database/sql"
	"fmt"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ProductRepositoryMysql struct {
	db *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		db: db,
	}
}

func (pr *ProductRepositoryMysql) FindById(id string) (*entity.Product, error) {
	var product entity.Product

	sqlStatement := `
		SELECT
			id,
			name,
			description,
			price,
			image,
			category_id,
			company_id
		FROM products
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := pr.db.QueryRow(sqlStatement, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.CategoryId,
		&product.CompanyId,
	)
	fmt.Println("erro no scan", err)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	sql := `SELECT * FROM products WHERE deleted_at IS NULL`
	rows, err := pr.db.Query(sql)
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
	sql := `INSERT INTO products (id, name, description, price, image, category_id, company_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := pr.db.Exec(
		sql,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Image,
		product.CategoryId,
		product.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryMysql) Update(product *entity.Product) error {
	sql := `UPDATE products SET name = $1, description = $2, price = $3, image = $4, category_id = $5, company_id = $6 WHERE id = $7`
	_, err := pr.db.Exec(
		sql,
		product.Name,
		product.Description,
		product.Price,
		product.Image,
		product.CategoryId,
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
	_, err := pr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryMysql) FindByCategoryId(categoryId string) ([]*entity.Product, error) {
	sql := `SELECT * FROM products WHERE category_id = $1 AND deleted_at IS NULL`
	rows, err := pr.db.Query(sql, categoryId)
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
