package entity

// Question
type Question struct {
	ID                 int    `json:"id" db:"id"`
	Text               string `json:"text" db:"text"`
	Number             int    `json:"number" db:"number"`
	Program            string `json:"program" db:"program"`
	Explanation        string `json:"explanation" db:"explanation"`
	QuestionCategoryID int    `json:"question_category_id" db:"question_category_id"`
}
