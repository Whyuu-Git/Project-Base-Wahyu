package service

import (
	"errors"

	"project-base-wahyu/internal/dto/response"
	"project-base-wahyu/internal/repository"
)

type ReportService interface {
	GetSummary(logID int) (*response.ExamSummaryResponse, error)
	GetExplanation(logID int) (*response.ExplanationResponse, error)
	GetRanking(tryoutCode string) (*response.RankingResponse, error)
	GetStudentReport(studentID int) (*response.StudentReportResponse, error)
}

type reportService struct {
	logExamRepo       repository.LogExamRepository
	detailLogRepo     repository.DetailLogRepository
	historyAnswerRepo repository.HistoryAnswerRepository
	studentRepo       repository.StudentRepository
	tryoutRepo        repository.TryoutRepository
}

func NewReportService(
	logExamRepo repository.LogExamRepository,
	detailLogRepo repository.DetailLogRepository,
	historyAnswerRepo repository.HistoryAnswerRepository,
	studentRepo repository.StudentRepository,
	tryoutRepo repository.TryoutRepository,
) ReportService {
	return &reportService{
		logExamRepo:       logExamRepo,
		detailLogRepo:     detailLogRepo,
		historyAnswerRepo: historyAnswerRepo,
		studentRepo:       studentRepo,
		tryoutRepo:        tryoutRepo,
	}
}

func (s *reportService) GetSummary(logID int) (*response.ExamSummaryResponse, error) {
	logExam, err := s.logExamRepo.GetByID(logID)
	if err != nil {
		return nil, errors.New("log ujian tidak ditemukan")
	}

	student, err := s.studentRepo.GetByID(logExam.StudentID)
	if err != nil {
		return nil, err
	}
	tryout, err := s.tryoutRepo.GetByID(logExam.TryoutCodeID)
	if err != nil {
		return nil, err
	}

	details, err := s.detailLogRepo.GetByLogID(logID)
	if err != nil {
		return nil, err
	}

	result := &response.ExamSummaryResponse{
		TryoutName:  tryout.Name,
		StudentName: student.Name,
		TotalScore:  logExam.TotalScore,
		PassStatus:  logExam.PassStatus,
	}
	for _, d := range details {
		result.ScorePerCategory = append(result.ScorePerCategory, response.CategoryScoreSummary{
			Name:       d.CategoryName,
			Score:      d.Score,
			PassStatus: d.PassStatus,
		})
	}

	return result, nil
}

func (s *reportService) GetExplanation(logID int) (*response.ExplanationResponse, error) {
	logExam, err := s.logExamRepo.GetByID(logID)
	if err != nil {
		return nil, errors.New("log ujian tidak ditemukan")
	}

	student, err := s.studentRepo.GetByID(logExam.StudentID)
	if err != nil {
		return nil, err
	}
	tryout, err := s.tryoutRepo.GetByID(logExam.TryoutCodeID)
	if err != nil {
		return nil, err
	}

	histories, err := s.historyAnswerRepo.GetByLogID(logID)
	if err != nil {
		return nil, err
	}

	result := &response.ExplanationResponse{
		Name:      student.Name,
		Email:     student.Email,
		Address:   student.Address,
		Phone:     student.Phone,
		ExamCode:  tryout.Code,
		StartDate: logExam.StartDate.Format("2006-01-02 15:04:05"),
		EndDate:   logExam.EndDate.Format("2006-01-02 15:04:05"),
	}

	for _, h := range histories {
		options, err := s.historyAnswerRepo.GetOptionsByHistoryAnswerID(h.ID)
		if err != nil {
			return nil, err
		}

		item := response.ExplanationQuestionItem{
			ID:          h.ID,
			Number:      h.Number,
			Text:        h.Question,
			Explanation: h.Explanations,
		}
		if h.AnswerID != nil {
			item.AnswerID = *h.AnswerID
			if selected, err := s.historyAnswerRepo.GetOptionByID(*h.AnswerID); err == nil {
				item.IsTrue = selected.IsTrue
			}
		}
		for _, opt := range options {
			item.Answers = append(item.Answers, response.ExplanationAnswerItem{
				ID:     opt.ID,
				Option: opt.Option,
				Text:   opt.Text,
				IsTrue: opt.IsTrue,
			})
		}
		result.StudentAnswers = append(result.StudentAnswers, item)
	}

	return result, nil
}

func (s *reportService) GetRanking(tryoutCode string) (*response.RankingResponse, error) {
	tryout, err := s.tryoutRepo.GetByCode(tryoutCode)
	if err != nil {
		return nil, errors.New("kode ujian tidak ditemukan")
	}

	rows, err := s.logExamRepo.GetRankingByTryoutCode(tryoutCode)
	if err != nil {
		return nil, err
	}

	result := &response.RankingResponse{TryoutName: tryout.Name}
	for _, row := range rows {
		result.Ranks = append(result.Ranks, response.RankingItem{
			StudentName: row.StudentName,
			TotalScore:  row.TotalScore,
			PassStatus:  row.PassStatus,
		})
	}

	return result, nil
}

func (s *reportService) GetStudentReport(studentID int) (*response.StudentReportResponse, error) {
	student, err := s.studentRepo.GetByID(studentID)
	if err != nil {
		return nil, errors.New("siswa tidak ditemukan")
	}

	rows, err := s.logExamRepo.GetReportByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	result := &response.StudentReportResponse{
		Name:    student.Name,
		Address: student.Address,
	}
	for _, row := range rows {
		result.Report = append(result.Report, response.StudentReportItem{
			TryoutName: row.TryoutName,
			TotalScore: row.TotalScore,
			PassStatus: row.PassStatus,
			Repeat:     row.Repeat,
		})
	}

	return result, nil
}
