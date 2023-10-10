package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type ProductRepositoryPostgresql struct {
	db *sql.DB
}

func NewProductRepositoryPostgresql(db *sql.DB) *ProductRepositoryPostgresql {
	return &ProductRepositoryPostgresql{
		db: db,
	}
}

func (pr *ProductRepositoryPostgresql) FindById(id string) (*entity.Product, error) {
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
		WHERE id = $1
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

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepositoryPostgresql) FindAll() ([]*entity.Product, error) {
	sql := `
		SELECT
			id,
			name,
			description,
			price,
			image,
			category_id,
			company_id 
		FROM products 
		WHERE deleted_at IS NULL
	`
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

func (pr *ProductRepositoryPostgresql) Create(product *entity.Product) error {
	sql := `
		INSERT INTO
			products (
				id,
				name,
				description,
				price,
				image,
				category_id,
				company_id,
				created_at,
				updated_at
			)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			NOW(),
			NOW()
		)
	`
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

func (pr *ProductRepositoryPostgresql) Update(product *entity.Product) error {
	sql := `
		UPDATE products
		SET
			name = $1,
			description = $2,
			price = $3,
			image = $4,
			category_id = $5,
			company_id = $6,
		WHERE
			id = $7
	`
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

func (pr *ProductRepositoryPostgresql) Delete(id string) error {
	sql := `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = $1
	`
	_, err := pr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
