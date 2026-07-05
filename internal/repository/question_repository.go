package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)

type QuestionRepository interface {
	GetByModuleID(moduleID int) ([]entity.Question, error)
	GetAnswersByQuestionID(questionID int) ([]entity.Answer, error)
	CreateQuestion(question *entity.Question) (int, error)
	CreateAnswer(answer *entity.Answer) (int, error)
}

type questionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) GetByModuleID(moduleID int) ([]entity.Question, error) {
	var questions []entity.Question
	query := `
		SELECT q.id, q.text, q.number, q.program, q.explanation, q.question_category_id
		FROM questions q
		JOIN module_questions mq ON mq.question_id = q.id
		WHERE mq.module_id = $1`

	err := r.db.Select(&questions, query, moduleID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *questionRepository) GetAnswersByQuestionID(questionID int) ([]entity.Answer, error) {
	var answers []entity.Answer
	query := `SELECT id, score, option, text, is_true, questions_id
	          FROM answers WHERE questions_id = $1`

	err := r.db.Select(&answers, query, questionID)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *questionRepository) CreateQuestion(question *entity.Question) (int, error) {
	var id int
	query := `
		INSERT INTO questions (text, number, program, explanation, question_category_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := r.db.QueryRow(query, question.Text, question.Number, question.Program, question.Explanation, question.QuestionCategoryID).Scan(&id)
	return id, err
}

func (r *questionRepository) CreateAnswer(answer *entity.Answer) (int, error) {
	var id int
	query := `
		INSERT INTO answers (score, option, text, is_true, questions_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := r.db.QueryRow(query, answer.Score, answer.Option, answer.Text, answer.IsTrue, answer.QuestionsID).Scan(&id)
	return id, err
}
