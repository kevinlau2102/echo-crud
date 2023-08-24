package repository

import (
	"crud_echo/pkg/domain"
	"database/sql"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) domain.StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (sr StudentRepository) GetStudents() ([]domain.Student, error) {
	sql := `SELECT * FROM student`
	rows, err := sr.db.Query(sql)
	var students []domain.Student
	for rows.Next() {
		var student domain.Student
		err2 := rows.Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)
		if err2 != nil {
			return nil, err2
		}
		students = append(students, student)
	}
	return students, err
}
