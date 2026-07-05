package entity

// Answer
type Answer struct {
	ID          int     `json:"id" db:"id"`
	Score       float64 `json:"score" db:"score"`
	Option      string  `json:"option" db:"option"`
	Text        string  `json:"text" db:"text"`
	IsTrue      bool    `json:"is_true" db:"is_true"`
	QuestionsID int     `json:"questions_id" db:"questions_id"`
}
