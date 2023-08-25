package repository

import (
	"crud_echo/pkg/domain"
	"database/sql"
)

type StudentRepository struct {
	db *sql.DB // nil
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

func (sr StudentRepository) GetStudent(id int) (domain.Student, error) {
	var student domain.Student
	sql := `SELECT * FROM student WHERE id = $1`

	err := sr.db.QueryRow(sql, id).Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)
	return student, err
}

func (sr StudentRepository) CreateStudent(req domain.Student) error {
	sql := `INSERT INTO student (fullname, address, birthdate, class, batch, school_name) values ($1, $2, $3, $4, $5, $6)`
	stmt, err := sr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(req.Fullname, req.Address, req.Birthdate, req.Class, req.Batch, req.SchoolName)
	if err2 != nil {
		return err2
	}
	return nil
}
