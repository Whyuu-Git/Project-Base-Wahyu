package repository

import (
	"github.com/jmoiron/sqlx"

	"project-base-wahyu/internal/entity"
)
type LogExamRepository interface {
	Create(logExam *entity.LogExam) (int, error)
	GetByID(id int) (*entity.LogExam, error)
	UpdateScore(id int, totalScore float64, passStatus bool) error
	GetRankingByTryoutCode(tryoutCode string) ([]RankingRow, error)
	GetReportByStudentID(studentID int) ([]ReportRow, error)
}

type RankingRow struct {
	LogID       int     `db:"log_id"`
	StudentName string  `db:"student_name"`
	TotalScore  float64 `db:"total_score"`
	PassStatus  bool    `db:"pass_status"`
}

type ReportRow struct {
	LogID      int     `db:"log_id"`
	TryoutName string  `db:"tryout_name"`
	TotalScore float64 `db:"total_score"`
	PassStatus bool    `db:"pass_status"`
	Repeat     int     `db:"repeat"`
}

type logExamRepository struct {
	db *sqlx.DB
}

func NewLogExamRepository(db *sqlx.DB) LogExamRepository {
	return &logExamRepository{db: db}
}

func (r *logExamRepository) Create(logExam *entity.LogExam) (int, error) {
	var id int
	query := `
		INSERT INTO log_exam (tryout_code_id, pass_status, total_score, repeat, start_date, student_id)
		VALUES ($1, $2, $3, $4, NOW(), $5)
		RETURNING id`

	err := r.db.QueryRow(
		query,
		logExam.TryoutCodeID,
		false,
		0,
		logExam.Repeat,
		logExam.StudentID,
	).Scan(&id)

	return id, err
}

func (r *logExamRepository) GetByID(id int) (*entity.LogExam, error) {
	var logExam entity.LogExam
	err := r.db.Get(&logExam, `SELECT * FROM log_exam WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &logExam, nil
}

func (r *logExamRepository) UpdateScore(id int, totalScore float64, passStatus bool) error {
	query := `UPDATE log_exam SET total_score = $1, pass_status = $2, end_date = NOW() WHERE id = $3`
	_, err := r.db.Exec(query, totalScore, passStatus, id)
	return err
}

func (r *logExamRepository) GetRankingByTryoutCode(tryoutCode string) ([]RankingRow, error) {
	var rows []RankingRow
	query := `
		SELECT le.id AS log_id, s.name AS student_name, le.total_score, le.pass_status
		FROM log_exam le
		JOIN tryout_codes tc ON tc.id = le.tryout_code_id
		JOIN student s ON s.id = le.student_id
		WHERE tc.code = $1
		ORDER BY le.total_score DESC`

	err := r.db.Select(&rows, query, tryoutCode)
	return rows, err
}

func (r *logExamRepository) GetReportByStudentID(studentID int) ([]ReportRow, error) {
	var rows []ReportRow
	query := `
		SELECT le.id AS log_id, tc.name AS tryout_name, le.total_score, le.pass_status, le.repeat
		FROM log_exam le
		JOIN tryout_codes tc ON tc.id = le.tryout_code_id
		WHERE le.student_id = $1
		ORDER BY le.id DESC`

	err := r.db.Select(&rows, query, studentID)
	return rows, err
}
