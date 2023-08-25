package controller

import (
	"crud_echo/pkg/domain"
	"crud_echo/pkg/dto"
	"crud_echo/shared/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type StudentControler struct {
	StudentUsecase domain.StudentUsecase
}

func (sc *StudentControler) GetStudents(c echo.Context) error {
	resp, err := sc.StudentUsecase.GetStudents()
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success", resp)
}

func (sc *StudentControler) GetStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := sc.StudentUsecase.GetStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "id student not found", nil)
	}
	return response.SetResponse(c, http.StatusOK, "success", resp)
}

func (sc *StudentControler) CreateStudent(c echo.Context) error {
	var studentdto dto.StudentDTO
	if err := c.Bind(&studentdto); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	if err := studentdto.Validation(); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	if err := sc.StudentUsecase.CreateStudent(studentdto); err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success", nil)
}
