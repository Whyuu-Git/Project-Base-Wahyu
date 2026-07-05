package router

import (
	"github.com/gofiber/fiber/v2"

	"project-base-wahyu/internal/handler"
)

func SetupRoutes(
	app *fiber.App,
	examHandler *handler.ExamHandler,
	studentHandler *handler.StudentHandler,
	reportHandler *handler.ReportHandler,
	moduleHandler *handler.ModuleHandler,
) {
	api := app.Group("/api/v1")

	exam := api.Group("/exam")
	exam.Get("/codes/:examCode", examHandler.GetExamByCode)    
	exam.Post("/answers", examHandler.SubmitAnswer)             
	exam.Get("/summary/:logId", examHandler.GetSummary)         
	exam.Get("/explanation/:logId", examHandler.GetExplanation) 

	students := api.Group("/students")
	students.Post("/register", studentHandler.Register)         
	students.Get("/:id/report", reportHandler.GetStudentReport) 

	tryout := api.Group("/tryout")
	tryout.Get("/:code/ranking", reportHandler.GetRanking) 

	
	admin := api.Group("/admin")
	admin.Post("/categories", moduleHandler.CreateCategory)
	admin.Post("/questions", moduleHandler.CreateQuestion)
	admin.Post("/modules", moduleHandler.CreateModule)
	admin.Post("/modules/:moduleId/questions", moduleHandler.AssignQuestion)
	admin.Post("/tryout-codes", moduleHandler.CreateTryoutCode)
}
