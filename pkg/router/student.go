package router

import (
	"crud_echo/pkg/controller"
	"crud_echo/pkg/repository"
	"crud_echo/pkg/usecase"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func NewStudentRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	sr := repository.NewStudentRepository(db)
	su := usecase.NewStudentUsecase(sr)
	sc := &controller.StudentControler{
		StudentUsecase: su,
	}

	e.GET("/student", sc.GetStudents)
	e.GET("/student/:id", sc.GetStudent)
	e.POST("/student", sc.CreateStudent)
	e.PUT("/student/:id", sc.UpdateStudent)
	e.DELETE("/student/:id", sc.DeleteStudent)
}
