package response

// ExplanationResponse
type ExplanationResponse struct {
	Name           string                    `json:"name"`
	Email          string                    `json:"email"`
	Address        string                    `json:"address"`
	Phone          string                    `json:"phone"`
	ExamCode       string                    `json:"exam_code"`
	StartDate      string                    `json:"start_date"`
	EndDate        string                    `json:"end_date"`
	StudentAnswers []ExplanationQuestionItem `json:"student_answers"`
}

// ExplanationQuestionItem .
type ExplanationQuestionItem struct {
	ID          int                     `json:"id"`
	Number      int                     `json:"number"`
	Text        string                  `json:"text"`
	Explanation string                  `json:"explanation"`
	AnswerID    int                     `json:"answer_id"`
	IsTrue      bool                    `json:"is_true"`
	Answers     []ExplanationAnswerItem `json:"answers"`
}

// ExplanationAnswerItem 
type ExplanationAnswerItem struct {
	ID     int    `json:"id"`
	Option string `json:"option"`
	Text   string `json:"text"`
	IsTrue bool   `json:"is_true"`
}
