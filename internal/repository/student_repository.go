package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)

type StudentRepository interface {
	GetByEmail(email string) (*entity.Student, error)
	GetByID(id int) (*entity.Student, error)
	Create(student *entity.Student) (int, error)
}

type studentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) GetByEmail(email string) (*entity.Student, error) {
	var student entity.Student
	err := r.db.Get(&student, `SELECT * FROM student WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) GetByID(id int) (*entity.Student, error) {
	var student entity.Student
	err := r.db.Get(&student, `SELECT * FROM student WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) Create(student *entity.Student) (int, error) {
	var id int
	query := `
		INSERT INTO student (name, email, address, phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := r.db.QueryRow(query, student.Name, student.Email, student.Address, student.Phone).Scan(&id)
	return id, err
}
