package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)


type DetailLogRepository interface {
	Create(detailLog *entity.DetailLog) error
	GetByLogID(logID int) ([]DetailLogWithCategory, error)
}


type DetailLogWithCategory struct {
	CategoryName string  `db:"category_name"`
	Score        float64 `db:"score"`
	PassStatus   bool    `db:"pass_status"`
}

type detailLogRepository struct {
	db *sqlx.DB
}


func NewDetailLogRepository(db *sqlx.DB) DetailLogRepository {
	return &detailLogRepository{db: db}
}

func (r *detailLogRepository) Create(detailLog *entity.DetailLog) error {
	query := `INSERT INTO detail_log (log_id, category_id, score, pass_status) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, detailLog.LogID, detailLog.CategoryID, detailLog.Score, detailLog.PassStatus)
	return err
}

func (r *detailLogRepository) GetByLogID(logID int) ([]DetailLogWithCategory, error) {
	var rows []DetailLogWithCategory
	query := `
		SELECT qc.name AS category_name, dl.score, dl.pass_status
		FROM detail_log dl
		JOIN question_categories qc ON qc.id = dl.category_id
		WHERE dl.log_id = $1`

	err := r.db.Select(&rows, query, logID)
	return rows, err
}
