package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type GroupRepositoryPostgresql struct {
	db *sql.DB
}

func NewGroupRepositoryPostgresql(db *sql.DB) *GroupRepositoryPostgresql {
	return &GroupRepositoryPostgresql{
		db: db,
	}
}

func (gr *GroupRepositoryPostgresql) FindById(id string) (*entity.Group, error) {
	var group entity.Group

	sqlStatement := `
		SELECT
			id,
			name,
			company_id
		FROM groups
		WHERE id = $1
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

func (gr *GroupRepositoryPostgresql) FindAll() ([]*entity.Group, error) {
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

func (gr *GroupRepositoryPostgresql) Create(group *entity.Group) error {
	sql := `
		INSERT INTO groups (
			id,
			name,
			company_id,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
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

func (gr *GroupRepositoryPostgresql) Update(group *entity.Group) error {
	sql := `
		UPDATE
			groups
		SET
			name = $1,
			company_id = $2,
			updated_at = NOW()
		WHERE
			id = $3
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

func (gr *GroupRepositoryPostgresql) Delete(id string) error {
	sql := `
		UPDATE
			groups
		SET deleted_at = NOW()
		WHERE id = $1
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
