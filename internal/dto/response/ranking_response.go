package response

// RankingResponse
type RankingResponse struct {
	TryoutName string        `json:"tryout_name"`
	Ranks      []RankingItem `json:"ranks"`
}

// RankingItem 
type RankingItem struct {
	StudentName string  `json:"student_name"`
	TotalScore  float64 `json:"total_score"`
	PassStatus  bool    `json:"pass_status"`
}
