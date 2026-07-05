package response

// ExamQuestionResponse
type ExamQuestionResponse struct {
	TryoutName  string             `json:"tryout_name"`
	Instruction string             `json:"instruction"`
	Questions   []ExamQuestionItem `json:"questions"`
}

// ExamQuestionItem 
type ExamQuestionItem struct {
	ID          int              `json:"id"`
	Number      int              `json:"number"`
	Text        string           `json:"text"`
	Explanation string           `json:"explanation"`
	Answers     []ExamAnswerItem `json:"answers"`
}

// ExamAnswerItem 
type ExamAnswerItem struct {
	ID     int    `json:"id"`
	Option string `json:"option"`
	Text   string `json:"text"`
}
