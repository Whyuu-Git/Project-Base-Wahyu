package response

// SubmitAnswerResponse
type SubmitAnswerResponse struct {
	Name           string             `json:"name"`
	Email          string             `json:"email"`
	Address        string             `json:"address"`
	Phone          string             `json:"phone"`
	ExamCode       string             `json:"exam_code"`
	StartDate      string             `json:"start_date"`
	EndDate        string             `json:"end_date"`
	StudentAnswers []SubmitAnswerItem `json:"student_answers"`
}

// SubmitAnswerItem 
type SubmitAnswerItem struct {
	ID       int              `json:"id"`
	Number   int              `json:"number"`
	Text     string           `json:"text"`
	AnswerID int              `json:"answer_id"`
	Answers  []ExamAnswerItem `json:"answers"`
}
