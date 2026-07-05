package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)


type HistoryAnswerRepository interface {

	CreateSnapshot(history *entity.HistoryAnswer, options []entity.AnswerQuestion) (int, error)
	GetByLogID(logID int) ([]entity.HistoryAnswer, error)
	GetByID(id int) (*entity.HistoryAnswer, error)
	GetOptionsByHistoryAnswerID(historyAnswerID int) ([]entity.AnswerQuestion, error)
	GetOptionByID(id int) (*entity.AnswerQuestion, error)
	SetSelectedAnswer(historyAnswerID int, answerQuestionID int) error
}

type historyAnswerRepository struct {
	db *sqlx.DB
}

func NewHistoryAnswerRepository(db *sqlx.DB) HistoryAnswerRepository {
	return &historyAnswerRepository{db: db}
}

func (r *historyAnswerRepository) CreateSnapshot(history *entity.HistoryAnswer, options []entity.AnswerQuestion) (int, error) {
	// Mulai transaksi baru.
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var historyID int
	insertHistory := `
		INSERT INTO history_answer (log_id, question_id, number, question, explanations, question_category_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err = tx.QueryRow(
		insertHistory,
		history.LogID,
		history.QuestionID,
		history.Number,
		history.Question,
		history.Explanations,
		history.QuestionCategoryID,
	).Scan(&historyID) 
	if err != nil {
		return 0, err
	}
	insertOption := `
		INSERT INTO answer_questions (history_answer_id, answer_id, option, text, is_true)
		VALUES ($1, $2, $3, $4, $5)`
	for _, opt := range options {
		_, err = tx.Exec(insertOption, historyID, opt.AnswerID, opt.Option, opt.Text, opt.IsTrue)
		if err != nil {
			return 0, err
		}
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return historyID, nil
}

func (r *historyAnswerRepository) GetByLogID(logID int) ([]entity.HistoryAnswer, error) {
	var histories []entity.HistoryAnswer
	query := `SELECT * FROM history_answer WHERE log_id = $1 ORDER BY number ASC`

	err := r.db.Select(&histories, query, logID)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *historyAnswerRepository) GetByID(id int) (*entity.HistoryAnswer, error) {
	var history entity.HistoryAnswer
	err := r.db.Get(&history, `SELECT * FROM history_answer WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *historyAnswerRepository) GetOptionByID(id int) (*entity.AnswerQuestion, error) {
	var option entity.AnswerQuestion
	err := r.db.Get(&option, `SELECT * FROM answer_questions WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *historyAnswerRepository) GetOptionsByHistoryAnswerID(historyAnswerID int) ([]entity.AnswerQuestion, error) {
	var options []entity.AnswerQuestion
	query := `SELECT * FROM answer_questions WHERE history_answer_id = $1 ORDER BY id ASC`

	err := r.db.Select(&options, query, historyAnswerID)
	if err != nil {
		return nil, err
	}
	return options, nil
}

func (r *historyAnswerRepository) SetSelectedAnswer(historyAnswerID int, answerQuestionID int) error {
	query := `UPDATE history_answer SET answer_id = $1 WHERE id = $2`
	_, err := r.db.Exec(query, answerQuestionID, historyAnswerID)
	return err
}
