package request

import "time"

// CreateCategoryRequest
type CreateCategoryRequest struct {
	Name         string `json:"name" validate:"required"`
	PassingGrade int    `json:"passing_grade" validate:"required"`
	Program      string `json:"program"`
}

// CreateQuestionRequest.
type CreateQuestionRequest struct {
	Text               string                `json:"text" validate:"required"`
	Number             int                   `json:"number"`
	Program            string                `json:"program"`
	Explanation        string                `json:"explanation"`
	QuestionCategoryID int                   `json:"question_category_id" validate:"required"`
	Answers            []CreateAnswerRequest `json:"answers" validate:"required,dive"`
}

// CreateAnswerRequest 
type CreateAnswerRequest struct {
	Score  float64 `json:"score"`
	Option string  `json:"option" validate:"required"`
	Text   string  `json:"text" validate:"required"`
	IsTrue bool    `json:"is_true"`
}

// CreateModuleRequest 
type CreateModuleRequest struct {
	Code    string `json:"code" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Program string `json:"program"`
}

// AssignQuestionRequest
type AssignQuestionRequest struct {
	QuestionID int `json:"question_id" validate:"required"`
}

// CreateTryoutCodeRequest
type CreateTryoutCodeRequest struct {
	Code        string    `json:"code" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	ModuleID    int       `json:"module_id" validate:"required"`
	Instruction string    `json:"instruction"`
}
