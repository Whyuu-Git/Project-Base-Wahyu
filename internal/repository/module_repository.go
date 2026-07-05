package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)

type ModuleRepository interface {
	Create(module *entity.Module) (int, error)
	GetByID(id int) (*entity.Module, error)
	AssignQuestion(moduleID int, questionID int) error
}

type moduleRepository struct {
	db *sqlx.DB
}

func NewModuleRepository(db *sqlx.DB) ModuleRepository {
	return &moduleRepository{db: db}
}

func (r *moduleRepository) Create(module *entity.Module) (int, error) {
	var id int
	query := `INSERT INTO module (code, name, program) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, module.Code, module.Name, module.Program).Scan(&id)
	return id, err
}

func (r *moduleRepository) GetByID(id int) (*entity.Module, error) {
	var module entity.Module
	err := r.db.Get(&module, `SELECT * FROM module WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *moduleRepository) AssignQuestion(moduleID int, questionID int) error {
	query := `INSERT INTO module_questions (module_id, question_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, moduleID, questionID)
	return err
}
