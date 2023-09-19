package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type CategoryRepositoryMysql struct {
	db *sql.DB
}

func NewCategoryRepositoryMysql(db *sql.DB) *CategoryRepositoryMysql {
	return &CategoryRepositoryMysql{
		db: db,
	}
}

func (cr *CategoryRepositoryMysql) FindById(id string) (*entity.Category, error) {
	var category entity.Category

	sqlStatement := `
		SELECT
			id,
			name,
			company_id
		FROM categories
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := cr.db.QueryRow(sqlStatement, id).Scan(
		&category.ID,
		&category.Name,
		&category.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (cr *CategoryRepositoryMysql) FindAll() ([]*entity.Category, error) {
	sql := `
		SELECT
			id,
			name,
			company_id
		FROM categories
		WHERE deleted_at IS NULL
	`

	rows, err := cr.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CompanyId,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepositoryMysql) Create(category *entity.Category) error {
	sql := `
		INSERT INTO categories (
			id,
			name,
			company_id
		) VALUES (?, ?, ?)
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		category.ID,
		category.Name,
		category.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepositoryMysql) Update(category *entity.Category) error {
	sql := `
		UPDATE
			categories
		SET
			name = ?,
			company_id = ?,
			updated_at = NOW()
		WHERE
			id = ?
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		category.Name,
		category.CompanyId,
		category.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepositoryMysql) Delete(id string) error {
	sql := `
		UPDATE
			categories
		SET
			deleted_at = NOW()
		WHERE
			id = ?
	`

	stmt, err := cr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryMysql) FindProductsByCategoryId(categoryId string) ([]*entity.Product, error) {
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
		WHERE category_id = ? 
			AND deleted_at IS NULL
	`
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
