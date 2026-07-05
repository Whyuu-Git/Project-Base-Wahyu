package handler

import (
	"github.com/gofiber/fiber/v2"

	"project-base-wahyu/internal/pkg/response"
	"project-base-wahyu/internal/service"
)

type ReportHandler struct {
	reportService service.ReportService
}

func NewReportHandler(reportService service.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) GetRanking(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return response.Error(c, fiber.StatusBadRequest, "kode ujian tidak boleh kosong")
	}

	result, err := h.reportService.GetRanking(code)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, result)
}

func (h *ReportHandler) GetStudentReport(c *fiber.Ctx) error {
	studentID, err := c.ParamsInt("id")
	if err != nil || studentID == 0 {
		return response.Error(c, fiber.StatusBadRequest, "id siswa tidak valid")
	}

	result, err := h.reportService.GetStudentReport(studentID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, result)
}
