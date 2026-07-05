package response

// StudentReportResponse
type StudentReportResponse struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Report  []StudentReportItem `json:"report"`
}

// StudentReportItem 
type StudentReportItem struct {
	TryoutName string  `json:"tryout_name"`
	TotalScore float64 `json:"total_score"`
	PassStatus bool    `json:"pass_status"`
	Repeat     int     `json:"repeat"`
}
