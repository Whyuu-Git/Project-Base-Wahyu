package entity

// DetailLog
type DetailLog struct {
	ID         int     `json:"id" db:"id"`
	LogID      int     `json:"log_id" db:"log_id"`
	CategoryID int     `json:"category_id" db:"category_id"`
	Score      float64 `json:"score" db:"score"`
	PassStatus bool    `json:"pass_status" db:"pass_status"`
}
