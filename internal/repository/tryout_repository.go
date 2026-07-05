package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)

type TryoutRepository interface {
	GetByCode(code string) (*entity.TryoutCode, error)
	GetByID(id int) (*entity.TryoutCode, error)
	Create(tryout *entity.TryoutCode) (int, error)
}

type tryoutRepository struct {
	db *sqlx.DB
}

func NewTryoutRepository(db *sqlx.DB) TryoutRepository {
	return &tryoutRepository{db: db}
}

func (r *tryoutRepository) GetByCode(code string) (*entity.TryoutCode, error) {
	var tryout entity.TryoutCode
	query := `SELECT id, code, name, start_date, end_date, module_id, instruction
	          FROM tryout_codes WHERE code = $1`

	err := r.db.Get(&tryout, query, code)
	if err != nil {
		return nil, err
	}
	return &tryout, nil
}

func (r *tryoutRepository) Create(tryout *entity.TryoutCode) (int, error) {
	var id int
	query := `
		INSERT INTO tryout_codes (code, name, start_date, end_date, module_id, instruction)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err := r.db.QueryRow(query, tryout.Code, tryout.Name, tryout.StartDate, tryout.EndDate, tryout.ModuleID, tryout.Instruction).Scan(&id)
	return id, err
}

func (r *tryoutRepository) GetByID(id int) (*entity.TryoutCode, error) {
	var tryout entity.TryoutCode
	err := r.db.Get(&tryout, `SELECT * FROM tryout_codes WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &tryout, nil
}
