package response

// ExamSummaryResponse
type ExamSummaryResponse struct {
	TryoutName       string                 `json:"tryout_name"`
	StudentName      string                 `json:"student_name"`
	TotalScore       float64                `json:"total_score"`
	PassStatus       bool                   `json:"pass_status"`
	ScorePerCategory []CategoryScoreSummary `json:"score_per_category"`
}

// CategoryScoreSummary 
type CategoryScoreSummary struct {
	Name       string  `json:"name"`
	Score      float64 `json:"score"`
	PassStatus bool    `json:"pass_status"`
}
