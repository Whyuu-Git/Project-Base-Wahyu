package entity

// AnswerQuestion
type AnswerQuestion struct {
	ID              int    `json:"id" db:"id"`
	HistoryAnswerID int    `json:"history_answer_id" db:"history_answer_id"`
	AnswerID        int    `json:"answer_id" db:"answer_id"` // referensi ke answers.id (master)
	Option          string `json:"option" db:"option"`
	Text            string `json:"text" db:"text"`
	IsTrue          bool   `json:"is_true" db:"is_true"`
}
