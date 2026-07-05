package entity

import "time"

// LogExam
type LogExam struct {
	ID           int       `json:"id" db:"id"`
	TryoutCodeID int       `json:"tryout_code_id" db:"tryout_code_id"`
	PassStatus   bool      `json:"pass_status" db:"pass_status"`
	TotalScore   float64   `json:"total_score" db:"total_score"`
	Repeat       int       `json:"repeat" db:"repeat"`
	StartDate    time.Time `json:"start_date" db:"start_date"`
	EndDate      time.Time `json:"end_date" db:"end_date"`
	StudentID    int       `json:"student_id" db:"student_id"`
}
