package handler

import (
	"github.com/gofiber/fiber/v2"

	"project-base-wahyu/internal/dto/request"
	"project-base-wahyu/internal/pkg/response"
	"project-base-wahyu/internal/pkg/validator"
	"project-base-wahyu/internal/service"
)


type ExamHandler struct {
	examService   service.ExamService
	answerService service.AnswerService
	reportService service.ReportService
}

func NewExamHandler(examService service.ExamService, answerService service.AnswerService, reportService service.ReportService) *ExamHandler {
	return &ExamHandler{examService: examService, answerService: answerService, reportService: reportService}
}

func (h *ExamHandler) GetExamByCode(c *fiber.Ctx) error {
	
	examCode := c.Params("examCode")

	studentID := c.QueryInt("student_id", 0)

	if examCode == "" {
		return response.Error(c, fiber.StatusBadRequest, "exam code tidak boleh kosong")
	}
	if studentID == 0 {
		return response.Error(c, fiber.StatusBadRequest, "student_id wajib diisi")
	}

	result, err := h.examService.StartExam(examCode, studentID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusOK, result)
}

func (h *ExamHandler) SubmitAnswer(c *fiber.Ctx) error {
	var req request.SubmitAnswerRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}
	if err := validator.Validate(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "validasi gagal: "+err.Error())
	}

	result, err := h.answerService.SubmitAnswers(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, result)
}
func (h *ExamHandler) GetSummary(c *fiber.Ctx) error {
	logID, err := c.ParamsInt("logId")
	if err != nil || logID == 0 {
		return response.Error(c, fiber.StatusBadRequest, "logId tidak valid")
	}

	result, err := h.reportService.GetSummary(logID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, result)
}

func (h *ExamHandler) GetExplanation(c *fiber.Ctx) error {
	logID, err := c.ParamsInt("logId")
	if err != nil || logID == 0 {
		return response.Error(c, fiber.StatusBadRequest, "logId tidak valid")
	}

	result, err := h.reportService.GetExplanation(logID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, result)
}
