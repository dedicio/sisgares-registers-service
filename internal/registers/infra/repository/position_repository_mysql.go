package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type PositionRepositoryMysql struct {
	db *sql.DB
}

func NewPositionRepositoryMysql(db *sql.DB) *PositionRepositoryMysql {
	return &PositionRepositoryMysql{
		db: db,
	}
}

func (pr *PositionRepositoryMysql) FindById(id string) (*entity.Position, error) {
	var position entity.Position

	sqlStatement := `
		SELECT
			id,
			name,
			description,
			group_id,
			company_id
		FROM positions
		WHERE id = ?
			AND deleted_at IS NULL
	`
	err := pr.db.QueryRow(sqlStatement, id).Scan(
		&position.ID,
		&position.Name,
		&position.Description,
		&position.GroupId,
		&position.CompanyId,
	)

	if err != nil {
		return nil, err
	}

	return &position, nil
}

func (pr *PositionRepositoryMysql) FindAll() ([]*entity.Position, error) {
	sql := `
		SELECT
			id,
			name,
			description,
			group_id,
			company_id
		FROM positions
		WHERE deleted_at IS NULL
	`

	rows, err := pr.db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	positions := []*entity.Position{}
	for rows.Next() {
		var position entity.Position
		err := rows.Scan(
			&position.ID,
			&position.Name,
			&position.Description,
			&position.GroupId,
			&position.CompanyId,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, &position)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func (pr *PositionRepositoryMysql) Create(position *entity.Position) error {
	sql := `
		INSERT INTO positions (
			id,
			name,
			description,
			group_id,
			company_id,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			NOW(),
			NOW()
		)
	`

	_, err := pr.db.Exec(
		sql,
		position.ID,
		position.Name,
		position.Description,
		position.GroupId,
		position.CompanyId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *PositionRepositoryMysql) Update(position *entity.Position) error {
	sql := `
		UPDATE positions
		SET
			name = ?,
			description = ?,
			group_id = ?,
			company_id = ?,
			updated_at = NOW()
		WHERE id = ?
	`

	_, err := pr.db.Exec(
		sql,
		position.Name,
		position.Description,
		position.GroupId,
		position.CompanyId,
		position.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *PositionRepositoryMysql) Delete(id string) error {
	sql := `
		UPDATE positions
		SET deleted_at = NOW()
		WHERE id = ?
	`

	_, err := pr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
