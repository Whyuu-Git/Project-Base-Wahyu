package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"project-base-wahyu/config"
	"project-base-wahyu/internal/handler"
	"project-base-wahyu/internal/pkg/database"
	"project-base-wahyu/internal/repository"
	"project-base-wahyu/internal/router"
	"project-base-wahyu/internal/service"
)

func main() {
	
	// 1. Load konfigurasi dari .env
	cfg := config.LoadConfig()

	// 2. Koneksi ke database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("gagal koneksi database: %v", err)
	}
	defer db.Close()

	// 3a. Repository layer
	tryoutRepo := repository.NewTryoutRepository(db)
	questionRepo := repository.NewQuestionRepository(db)
	logExamRepo := repository.NewLogExamRepository(db)
	historyAnswerRepo := repository.NewHistoryAnswerRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	detailLogRepo := repository.NewDetailLogRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	moduleRepo := repository.NewModuleRepository(db)

	// 3b. Service layer 
	examService := service.NewExamService(tryoutRepo, questionRepo, logExamRepo, historyAnswerRepo)
	answerService := service.NewAnswerService(historyAnswerRepo, logExamRepo, detailLogRepo, categoryRepo, studentRepo, tryoutRepo)
	reportService := service.NewReportService(logExamRepo, detailLogRepo, historyAnswerRepo, studentRepo, tryoutRepo)
	studentService := service.NewStudentService(studentRepo)
	moduleService := service.NewModuleService(categoryRepo, questionRepo, moduleRepo, tryoutRepo)

	// 3c. Handler layer 
	examHandler := handler.NewExamHandler(examService, answerService, reportService)
	studentHandler := handler.NewStudentHandler(studentService)
	reportHandler := handler.NewReportHandler(reportService)
	moduleHandler := handler.NewModuleHandler(moduleService)

	// 4. Setup Fiber app & daftarkan semua route
	app := fiber.New()
	router.SetupRoutes(app, examHandler, studentHandler, reportHandler, moduleHandler)

	// 5. Jalankan server
	log.Printf("server berjalan di port %s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("gagal menjalankan server: %v", err)
	}
}
