package entity

// HistoryAnswer
type HistoryAnswer struct {
	ID                 int    `json:"id" db:"id"`
	LogID              int    `json:"log_id" db:"log_id"`
	QuestionID         int    `json:"question_id" db:"question_id"`
	AnswerID           *int   `json:"answer_id" db:"answer_id"` 
	Number             int    `json:"number" db:"number"`
	Question           string `json:"question" db:"question"`
	Explanations       string `json:"explanations" db:"explanations"`
	QuestionCategoryID int    `json:"question_category_id" db:"question_category_id"`
}
