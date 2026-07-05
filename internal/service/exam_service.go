package service

import (
	"errors"
	"math/rand"
	"time"

	"project-base-wahyu/internal/dto/response"
	"project-base-wahyu/internal/entity"
	"project-base-wahyu/internal/repository"
)

type ExamService interface {
	StartExam(tryoutCode string, studentID int) (*response.ExamQuestionResponse, error)
}

type examService struct {
	tryoutRepo        repository.TryoutRepository
	questionRepo      repository.QuestionRepository
	logExamRepo       repository.LogExamRepository
	historyAnswerRepo repository.HistoryAnswerRepository
}

func NewExamService(
	tryoutRepo repository.TryoutRepository,
	questionRepo repository.QuestionRepository,
	logExamRepo repository.LogExamRepository,
	historyAnswerRepo repository.HistoryAnswerRepository,
) ExamService {
	return &examService{
		tryoutRepo:        tryoutRepo,
		questionRepo:      questionRepo,
		logExamRepo:       logExamRepo,
		historyAnswerRepo: historyAnswerRepo,
	}
}

func (s *examService) StartExam(tryoutCode string, studentID int) (*response.ExamQuestionResponse, error) {
	tryout, err := s.tryoutRepo.GetByCode(tryoutCode)
	if err != nil {
		return nil, errors.New("kode ujian tidak ditemukan")
	}

	now := time.Now()
	if now.Before(tryout.StartDate) || now.After(tryout.EndDate) {
		return nil, errors.New("kode ujian belum aktif atau sudah berakhir")
	}

	questions, err := s.questionRepo.GetByModuleID(tryout.ModuleID)
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.New("modul belum memiliki soal")
	}

	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	logExamID, err := s.logExamRepo.Create(&entity.LogExam{
		TryoutCodeID: tryout.ID,
		StudentID:    studentID,
		Repeat:       1, 
	})
	if err != nil {
		return nil, err
	}

	examResponse := &response.ExamQuestionResponse{
		TryoutName:  tryout.Name,
		Instruction: tryout.Instruction,
	}

	for i, q := range questions {
		answers, err := s.questionRepo.GetAnswersByQuestionID(q.ID)
		if err != nil {
			return nil, err
		}
		rand.Shuffle(len(answers), func(a, b int) {
			answers[a], answers[b] = answers[b], answers[a]
		})

		historyAnswer := &entity.HistoryAnswer{
			LogID:              logExamID,
			QuestionID:         q.ID, 
			Number:             i + 1,
			Question:           q.Text,
			Explanations:       q.Explanation,
			QuestionCategoryID: q.QuestionCategoryID,
		}

		options := make([]entity.AnswerQuestion, 0, len(answers))
		for _, a := range answers {
			options = append(options, entity.AnswerQuestion{
				AnswerID: a.ID, 
				Option:   a.Option,
				Text:     a.Text,
				IsTrue:   a.IsTrue,
			})
		}

		historyAnswerID, err := s.historyAnswerRepo.CreateSnapshot(historyAnswer, options)
		if err != nil {
			return nil, err
		}
		savedOptions, err := s.historyAnswerRepo.GetOptionsByHistoryAnswerID(historyAnswerID)
		if err != nil {
			return nil, err
		}
		questionItem := response.ExamQuestionItem{
			ID:          historyAnswerID,
			Number:      i + 1,
			Text:        q.Text,
			Explanation: q.Explanation,
		}
		for _, opt := range savedOptions {
			questionItem.Answers = append(questionItem.Answers, response.ExamAnswerItem{
				ID:     opt.ID, 
				Option: opt.Option,
				Text:   opt.Text,
			})
		}

		examResponse.Questions = append(examResponse.Questions, questionItem)
	}
	return examResponse, nil
}
