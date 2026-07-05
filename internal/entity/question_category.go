package entity

// QuestionCategory
type QuestionCategory struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	PassingGrade int    `json:"passing_grade" db:"passing_grade"`
	Program      string `json:"program" db:"program"`
}
