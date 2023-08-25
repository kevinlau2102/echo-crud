package domain

import "crud_echo/pkg/dto"

type Student struct {
	Id         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

type StudentRepository interface {
	GetStudents() ([]Student, error)
	GetStudent(id int) (Student, error)
	CreateStudent(req Student) error
	UpdateStudent(req Student) error
	DeleteStudent(id int) error
}

type StudentUsecase interface {
	GetStudents() ([]Student, error)
	GetStudent(id int) (Student, error)
	CreateStudent(req dto.StudentDTO) error
	UpdateStudent(req dto.StudentDTO, id int) error
	DeleteStudent(id int) error
}
