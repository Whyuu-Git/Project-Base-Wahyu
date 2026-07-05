package handler

import (
	"github.com/gofiber/fiber/v2"

	"project-base-wahyu/internal/dto/request"
	"project-base-wahyu/internal/pkg/response"
	"project-base-wahyu/internal/pkg/validator"
	"project-base-wahyu/internal/service"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

func (h *StudentHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}
	
	if err := validator.Validate(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "validasi gagal: "+err.Error())
	}

	student, err := h.studentService.Register(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, student)
}
