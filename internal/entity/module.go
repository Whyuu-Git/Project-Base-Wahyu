package entity

// Module
type Module struct {
	ID      int    `json:"id" db:"id"`
	Code    string `json:"code" db:"code"`
	Name    string `json:"name" db:"name"`
	Program string `json:"program" db:"program"`
}
