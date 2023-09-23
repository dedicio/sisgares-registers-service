package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type GroupRepositoryMysql struct {
	db *sql.DB
}

func NewGroupRepositoryMysql(db *sql.DB) *GroupRepositoryMysql {
	return &GroupRepositoryMysql{
		db: db,
	}
}

func (gr *GroupRepositoryMysql) FindById(id string) (*entity.Group, error) {
	var group entity.Group

	sqlStatement := `
		SELECT
			id,
			name,
			company_id
		FROM groups
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := gr.db.QueryRow(sqlStatement, id).Scan(
		&group.ID,
		&group.Name,
		&group.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (gr *GroupRepositoryMysql) FindAll() ([]*entity.Group, error) {
	sql := `
		SELECT
			id,
			name,
			company_id
		FROM groups
		WHERE deleted_at IS NULL
	`

	rows, err := gr.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*entity.Group

	for rows.Next() {
		var group entity.Group

		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.CompanyId,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, &group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (gr *GroupRepositoryMysql) Create(group *entity.Group) error {
	sql := `
		INSERT INTO groups (
			id,
			name,
			company_id,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
			?,
			NOW(),
			NOW()
		)
	`

	stmt, err := gr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		group.ID,
		group.Name,
		group.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepositoryMysql) Update(group *entity.Group) error {
	sql := `
		UPDATE
			groups
		SET
			name = ?,
			company_id = ?,
			updated_at = NOW()
		WHERE
			id = ?
	`

	stmt, err := gr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		group.Name,
		group.CompanyId,
		group.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepositoryMysql) Delete(id string) error {
	sql := `
		UPDATE
			groups
		SET deleted_at = NOW()
		WHERE id = ?
	`

	stmt, err := gr.db.Prepare(sql)
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
