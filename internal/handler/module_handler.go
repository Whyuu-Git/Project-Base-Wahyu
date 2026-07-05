package handler

import (
	"github.com/gofiber/fiber/v2"

	"project-base-wahyu/internal/dto/request"
	"project-base-wahyu/internal/pkg/response"
	"project-base-wahyu/internal/service"
)

type ModuleHandler struct {
	moduleService service.ModuleService
}

func NewModuleHandler(moduleService service.ModuleService) *ModuleHandler {
	return &ModuleHandler{moduleService: moduleService}
}

func (h *ModuleHandler) CreateCategory(c *fiber.Ctx) error {
	var req request.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}
	result, err := h.moduleService.CreateCategory(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusCreated, result)
}

func (h *ModuleHandler) CreateQuestion(c *fiber.Ctx) error {
	var req request.CreateQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}
	result, err := h.moduleService.CreateQuestion(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusCreated, result)
}

func (h *ModuleHandler) CreateModule(c *fiber.Ctx) error {
	var req request.CreateModuleRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}
	result, err := h.moduleService.CreateModule(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusCreated, result)
}

func (h *ModuleHandler) AssignQuestion(c *fiber.Ctx) error {
	moduleID, err := c.ParamsInt("moduleId")
	if err != nil || moduleID == 0 {
		return response.Error(c, fiber.StatusBadRequest, "moduleId tidak valid")
	}

	var req request.AssignQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}

	if err := h.moduleService.AssignQuestionToModule(moduleID, req.QuestionID); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, fiber.Map{"message": "soal berhasil ditambahkan ke modul"})
}

func (h *ModuleHandler) CreateTryoutCode(c *fiber.Ctx) error {
	var req request.CreateTryoutCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format request tidak valid")
	}
	result, err := h.moduleService.CreateTryoutCode(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, fiber.StatusCreated, result)
}
