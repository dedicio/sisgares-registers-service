package repository

import (
	"database/sql"

	"github.com/dedicio/sisgares-registers-service/internal/registers/entity"
)

type PositionRepositoryPostgresql struct {
	db *sql.DB
}

func NewPositionRepositoryPostgresql(db *sql.DB) *PositionRepositoryPostgresql {
	return &PositionRepositoryPostgresql{
		db: db,
	}
}

func (pr *PositionRepositoryPostgresql) FindById(id string) (*entity.Position, error) {
	var position entity.Position

	sqlStatement := `
		SELECT
			id,
			name,
			description,
			group_id,
			company_id
		FROM positions
		WHERE id = $1
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

func (pr *PositionRepositoryPostgresql) FindAll(companyID string) ([]*entity.Position, error) {
	sql := `
		SELECT
			id,
			name,
			description,
			group_id,
			company_id
		FROM positions
		WHERE company_id = $1
			AND deleted_at IS NULL
	`

	rows, err := pr.db.Query(sql, companyID)

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

func (pr *PositionRepositoryPostgresql) Create(position *entity.Position) error {
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
			$1,
			$2,
			$3,
			$4,
			$5,
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

func (pr *PositionRepositoryPostgresql) Update(position *entity.Position) error {
	sql := `
		UPDATE positions
		SET
			name = $1,
			description = $2,
			group_id = $3,
			company_id = $4,
			updated_at = NOW()
		WHERE id = $5
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

func (pr *PositionRepositoryPostgresql) Delete(id string) error {
	sql := `
		UPDATE positions
		SET deleted_at = NOW()
		WHERE id = $1
	`

	_, err := pr.db.Exec(sql, id)

	if err != nil {
		return err
	}

	return nil
}
