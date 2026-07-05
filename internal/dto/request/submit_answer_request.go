package request

// SubmitAnswerRequest
type SubmitAnswerRequest struct {
	LogID   int                  `json:"log_id" validate:"required"`
	Answers []SubmitAnswerDetail `json:"answers" validate:"required,dive"`
}

// SubmitAnswerDetail 
type SubmitAnswerDetail struct {
	HistoryAnswerID  int `json:"history_answer_id" validate:"required"`
	AnswerQuestionID int `json:"answer_question_id" validate:"required"`
}
