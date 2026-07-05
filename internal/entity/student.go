package entity

// Student
type Student struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Email   string `json:"email" db:"email"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone"`
}
