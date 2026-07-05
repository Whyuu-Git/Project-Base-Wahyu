package service

import (
	"errors"

	"project-base-wahyu/internal/dto/request"
	"project-base-wahyu/internal/dto/response"
	"project-base-wahyu/internal/entity"
	"project-base-wahyu/internal/repository"
)

type AnswerService interface {
	SubmitAnswers(req request.SubmitAnswerRequest) (*response.SubmitAnswerResponse, error)
}

type answerService struct {
	historyAnswerRepo repository.HistoryAnswerRepository
	logExamRepo       repository.LogExamRepository
	detailLogRepo     repository.DetailLogRepository
	categoryRepo      repository.CategoryRepository
	studentRepo       repository.StudentRepository
	tryoutRepo        repository.TryoutRepository
}

func NewAnswerService(
	historyAnswerRepo repository.HistoryAnswerRepository,
	logExamRepo repository.LogExamRepository,
	detailLogRepo repository.DetailLogRepository,
	categoryRepo repository.CategoryRepository,
	studentRepo repository.StudentRepository,
	tryoutRepo repository.TryoutRepository,
) AnswerService {
	return &answerService{
		historyAnswerRepo: historyAnswerRepo,
		logExamRepo:       logExamRepo,
		detailLogRepo:     detailLogRepo,
		categoryRepo:      categoryRepo,
		studentRepo:       studentRepo,
		tryoutRepo:        tryoutRepo,
	}
}

func (s *answerService) SubmitAnswers(req request.SubmitAnswerRequest) (*response.SubmitAnswerResponse, error) {
	logExam, err := s.logExamRepo.GetByID(req.LogID)
	if err != nil {
		return nil, errors.New("log ujian tidak ditemukan")
	}

	for _, ans := range req.Answers {
		history, err := s.historyAnswerRepo.GetByID(ans.HistoryAnswerID)
		if err != nil {
			return nil, errors.New("soal tidak ditemukan")
		}
		if history.LogID != req.LogID {
			return nil, errors.New("soal tidak sesuai dengan sesi ujian ini")
		}

		option, err := s.historyAnswerRepo.GetOptionByID(ans.AnswerQuestionID)
		if err != nil {
			return nil, errors.New("opsi jawaban tidak ditemukan")
		}
		
		if option.HistoryAnswerID != ans.HistoryAnswerID {
			return nil, errors.New("opsi jawaban tidak sesuai dengan soal ini")
		}

		if err := s.historyAnswerRepo.SetSelectedAnswer(ans.HistoryAnswerID, ans.AnswerQuestionID); err != nil {
			return nil, err
		}
	}

	histories, err := s.historyAnswerRepo.GetByLogID(req.LogID)
	if err != nil {
		return nil, err
	}

	categoryCorrect := map[int]int{}
	categoryTotal := map[int]int{}

	for _, h := range histories {
		categoryTotal[h.QuestionCategoryID]++

		if h.AnswerID == nil {
			continue 
		}
		option, err := s.historyAnswerRepo.GetOptionByID(*h.AnswerID)
		if err != nil {
			continue
		}
		if option.IsTrue {
			categoryCorrect[h.QuestionCategoryID]++
		}
	}

	var totalScore float64
	overallPass := true

	for categoryID, total := range categoryTotal {
		correct := categoryCorrect[categoryID]

		score := (float64(correct) / float64(total)) * 100

		category, err := s.categoryRepo.GetByID(categoryID)
		if err != nil {
			return nil, err
		}

		passStatus := score >= float64(category.PassingGrade)
		if !passStatus {
			overallPass = false
		}

		if err := s.detailLogRepo.Create(&entity.DetailLog{
			LogID:      req.LogID,
			CategoryID: categoryID,
			Score:      score,
			PassStatus: passStatus,
		}); err != nil {
			return nil, err
		}

		totalScore += score
	}

	if len(categoryTotal) > 0 {
		totalScore = totalScore / float64(len(categoryTotal))
	}

	if err := s.logExamRepo.UpdateScore(req.LogID, totalScore, overallPass); err != nil {
		return nil, err
	}

	student, err := s.studentRepo.GetByID(logExam.StudentID)
	if err != nil {
		return nil, err
	}
	tryout, err := s.tryoutRepo.GetByID(logExam.TryoutCodeID)
	if err != nil {
		return nil, err
	}

	result := &response.SubmitAnswerResponse{
		Name:      student.Name,
		Email:     student.Email,
		Address:   student.Address,
		Phone:     student.Phone,
		ExamCode:  tryout.Code,
		StartDate: logExam.StartDate.Format("2006-01-02 15:04:05"),
	}

	for _, h := range histories {
		options, err := s.historyAnswerRepo.GetOptionsByHistoryAnswerID(h.ID)
		if err != nil {
			return nil, err
		}

		item := response.SubmitAnswerItem{
			ID:     h.ID,
			Number: h.Number,
			Text:   h.Question,
		}
		if h.AnswerID != nil {
			item.AnswerID = *h.AnswerID
		}
		for _, opt := range options {
			item.Answers = append(item.Answers, response.ExamAnswerItem{
				ID:     opt.ID,
				Option: opt.Option,
				Text:   opt.Text,
			})
		}
		result.StudentAnswers = append(result.StudentAnswers, item)
	}

	return result, nil
}
