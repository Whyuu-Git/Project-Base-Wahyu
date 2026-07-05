package service

import (
	"project-base-wahyu/internal/dto/request"
	"project-base-wahyu/internal/entity"
	"project-base-wahyu/internal/repository"
)


type StudentService interface {
	Register(req request.RegisterStudentRequest) (*entity.Student, error)
}

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{studentRepo: studentRepo}
}

func (s *studentService) Register(req request.RegisterStudentRequest) (*entity.Student, error) {
	existing, err := s.studentRepo.GetByEmail(req.Email)
	if err == nil && existing != nil {
		return existing, nil
	}

	student := &entity.Student{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}

	id, err := s.studentRepo.Create(student)
	if err != nil {
		return nil, err
	}
	student.ID = id

	return student, nil
}
