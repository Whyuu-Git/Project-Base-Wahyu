package request

// RegisterStudentRequest
type RegisterStudentRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
