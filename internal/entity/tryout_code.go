package entity

import "time"

// TryoutCode
type TryoutCode struct {
	ID          int       `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Name        string    `json:"name" db:"name"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	EndDate     time.Time `json:"end_date" db:"end_date"`
	ModuleID    int       `json:"module_id" db:"module_id"`
	Instruction string    `json:"instruction" db:"instruction"`
}
