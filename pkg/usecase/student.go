package usecase

import "crud_echo/pkg/domain"

type StudentUsecase struct {
	StudentRepository domain.StudentRepository
}

func NewStudentUsecase(studentRepository domain.StudentRepository) domain.StudentUsecase {
	return &StudentUsecase{
		StudentRepository: studentRepository,
	}
}

func (su StudentUsecase) GetStudents() ([]domain.Student, error) {
	return su.StudentRepository.GetStudents()
}
