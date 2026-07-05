package entity

// ModuleQuestion
type ModuleQuestion struct {
	ID         int `json:"id" db:"id"`
	ModuleID   int `json:"module_id" db:"module_id"`
	QuestionID int `json:"question_id" db:"question_id"`
}
