package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)


type CategoryRepository interface {
	GetByID(id int) (*entity.QuestionCategory, error)
	Create(category *entity.QuestionCategory) (int, error)
	GetAll() ([]entity.QuestionCategory, error)
}

type categoryRepository struct {
	db *sqlx.DB
}


func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetByID(id int) (*entity.QuestionCategory, error) {
	var category entity.QuestionCategory
	err := r.db.Get(&category, `SELECT * FROM question_categories WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *entity.QuestionCategory) (int, error) {
	var id int
	query := `INSERT INTO question_categories (name, passing_grade, program) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, category.Name, category.PassingGrade, category.Program).Scan(&id)
	return id, err
}

func (r *categoryRepository) GetAll() ([]entity.QuestionCategory, error) {
	var categories []entity.QuestionCategory
	err := r.db.Select(&categories, `SELECT * FROM question_categories ORDER BY id ASC`)
	return categories, err
}
