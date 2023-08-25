package usecase

import (
	"crud_echo/pkg/domain"
	"crud_echo/pkg/dto"
	"github.com/mitchellh/mapstructure"
)

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

func (su StudentUsecase) CreateStudent(req dto.StudentDTO) error {
	var student domain.Student
	mapstructure.Decode(req, &student)
	return su.StudentRepository.CreateStudent(student)
}

func (su StudentUsecase) GetStudent(id int) (domain.Student, error) {
	return su.StudentRepository.GetStudent(id)
}

func (su StudentUsecase) UpdateStudent(req dto.StudentDTO, id int) error {
	var student domain.Student
	mapstructure.Decode(req, &student)
	student.Id = id
	return su.StudentRepository.UpdateStudent(student)
}

func (su StudentUsecase) DeleteStudent(id int) error {
	return su.StudentRepository.DeleteStudent(id)
}
