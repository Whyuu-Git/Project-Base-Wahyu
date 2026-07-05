package service

import (
	"testing"
	"time"

	"project-base-wahyu/internal/entity"
	"project-base-wahyu/internal/repository"
)

type fakeTryoutRepo struct {
	tryout *entity.TryoutCode
	err    error
}

func (f *fakeTryoutRepo) GetByCode(code string) (*entity.TryoutCode, error) {
	return f.tryout, f.err
}
func (f *fakeTryoutRepo) GetByID(id int) (*entity.TryoutCode, error) { return f.tryout, f.err }
func (f *fakeTryoutRepo) Create(t *entity.TryoutCode) (int, error)   { return 0, nil }

type fakeQuestionRepo struct{}

func (f *fakeQuestionRepo) GetByModuleID(moduleID int) ([]entity.Question, error) { return nil, nil }
func (f *fakeQuestionRepo) GetAnswersByQuestionID(questionID int) ([]entity.Answer, error) {
	return nil, nil
}
func (f *fakeQuestionRepo) CreateQuestion(q *entity.Question) (int, error) { return 0, nil }
func (f *fakeQuestionRepo) CreateAnswer(a *entity.Answer) (int, error)     { return 0, nil }

type fakeLogExamRepo struct{}

func (f *fakeLogExamRepo) Create(l *entity.LogExam) (int, error)              { return 1, nil }
func (f *fakeLogExamRepo) GetByID(id int) (*entity.LogExam, error)            { return &entity.LogExam{}, nil }
func (f *fakeLogExamRepo) UpdateScore(id int, score float64, pass bool) error { return nil }
func (f *fakeLogExamRepo) GetRankingByTryoutCode(code string) ([]repository.RankingRow, error) {
	return nil, nil
}
func (f *fakeLogExamRepo) GetReportByStudentID(id int) ([]repository.ReportRow, error) {
	return nil, nil
}

type fakeHistoryAnswerRepo struct{}

func (f *fakeHistoryAnswerRepo) CreateSnapshot(h *entity.HistoryAnswer, o []entity.AnswerQuestion) (int, error) {
	return 1, nil
}
func (f *fakeHistoryAnswerRepo) GetByLogID(logID int) ([]entity.HistoryAnswer, error) {
	return nil, nil
}
func (f *fakeHistoryAnswerRepo) GetByID(id int) (*entity.HistoryAnswer, error) {
	return &entity.HistoryAnswer{}, nil
}
func (f *fakeHistoryAnswerRepo) GetOptionsByHistoryAnswerID(id int) ([]entity.AnswerQuestion, error) {
	return nil, nil
}
func (f *fakeHistoryAnswerRepo) GetOptionByID(id int) (*entity.AnswerQuestion, error) {
	return &entity.AnswerQuestion{}, nil
}
func (f *fakeHistoryAnswerRepo) SetSelectedAnswer(historyAnswerID int, answerQuestionID int) error {
	return nil
}

func TestStartExam_TryoutTidakDitemukan(t *testing.T) {
	svc := NewExamService(
		&fakeTryoutRepo{tryout: nil, err: errNotFound},
		&fakeQuestionRepo{},
		&fakeLogExamRepo{},
		&fakeHistoryAnswerRepo{},
	)

	_, err := svc.StartExam("KODE_TIDAK_ADA", 1)
	if err == nil {
		t.Fatal("seharusnya mengembalikan error karena tryout code tidak ditemukan, tapi tidak ada error")
	}
}

func TestStartExam_SudahBerakhir(t *testing.T) {
	expired := &entity.TryoutCode{
		ID:        1,
		ModuleID:  1,
		StartDate: time.Now().Add(-48 * time.Hour),
		EndDate:   time.Now().Add(-24 * time.Hour), 
	}

	svc := NewExamService(
		&fakeTryoutRepo{tryout: expired, err: nil},
		&fakeQuestionRepo{},
		&fakeLogExamRepo{},
		&fakeHistoryAnswerRepo{},
	)

	_, err := svc.StartExam("EXPIRED", 1)
	if err == nil {
		t.Fatal("seharusnya mengembalikan error karena tryout code sudah berakhir, tapi tidak ada error")
	}
}

var errNotFound = &notFoundError{}

type notFoundError struct{}

func (e *notFoundError) Error() string { return "not found" }
