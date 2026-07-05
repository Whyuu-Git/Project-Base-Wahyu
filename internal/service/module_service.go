package service

import (
	"project-base-wahyu/internal/dto/request"
	"project-base-wahyu/internal/entity"
	"project-base-wahyu/internal/repository"
)


type ModuleService interface {
	CreateCategory(req request.CreateCategoryRequest) (*entity.QuestionCategory, error)
	CreateQuestion(req request.CreateQuestionRequest) (*entity.Question, error)
	CreateModule(req request.CreateModuleRequest) (*entity.Module, error)
	AssignQuestionToModule(moduleID int, questionID int) error
	CreateTryoutCode(req request.CreateTryoutCodeRequest) (*entity.TryoutCode, error)
}

type moduleService struct {
	categoryRepo repository.CategoryRepository
	questionRepo repository.QuestionRepository
	moduleRepo   repository.ModuleRepository
	tryoutRepo   repository.TryoutRepository
}

func NewModuleService(
	categoryRepo repository.CategoryRepository,
	questionRepo repository.QuestionRepository,
	moduleRepo repository.ModuleRepository,
	tryoutRepo repository.TryoutRepository,
) ModuleService {
	return &moduleService{
		categoryRepo: categoryRepo,
		questionRepo: questionRepo,
		moduleRepo:   moduleRepo,
		tryoutRepo:   tryoutRepo,
	}
}

func (s *moduleService) CreateCategory(req request.CreateCategoryRequest) (*entity.QuestionCategory, error) {
	category := &entity.QuestionCategory{
		Name:         req.Name,
		PassingGrade: req.PassingGrade,
		Program:      req.Program,
	}
	id, err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}
	category.ID = id
	return category, nil
}

func (s *moduleService) CreateQuestion(req request.CreateQuestionRequest) (*entity.Question, error) {
	question := &entity.Question{
		Text:               req.Text,
		Number:             req.Number,
		Program:            req.Program,
		Explanation:        req.Explanation,
		QuestionCategoryID: req.QuestionCategoryID,
	}
	questionID, err := s.questionRepo.CreateQuestion(question)
	if err != nil {
		return nil, err
	}
	question.ID = questionID

	for _, a := range req.Answers {
		_, err := s.questionRepo.CreateAnswer(&entity.Answer{
			Score:       a.Score,
			Option:      a.Option,
			Text:        a.Text,
			IsTrue:      a.IsTrue,
			QuestionsID: questionID,
		})
		if err != nil {
			return nil, err
		}
	}

	return question, nil
}

func (s *moduleService) CreateModule(req request.CreateModuleRequest) (*entity.Module, error) {
	module := &entity.Module{
		Code:    req.Code,
		Name:    req.Name,
		Program: req.Program,
	}
	id, err := s.moduleRepo.Create(module)
	if err != nil {
		return nil, err
	}
	module.ID = id
	return module, nil
}

func (s *moduleService) AssignQuestionToModule(moduleID int, questionID int) error {
	return s.moduleRepo.AssignQuestion(moduleID, questionID)
}

func (s *moduleService) CreateTryoutCode(req request.CreateTryoutCodeRequest) (*entity.TryoutCode, error) {
	tryout := &entity.TryoutCode{
		Code:        req.Code,
		Name:        req.Name,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		ModuleID:    req.ModuleID,
		Instruction: req.Instruction,
	}
	id, err := s.tryoutRepo.Create(tryout)
	if err != nil {
		return nil, err
	}
	tryout.ID = id
	return tryout, nil
}
