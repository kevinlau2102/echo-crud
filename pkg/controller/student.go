package controller

import (
	"crud_echo/pkg/domain"
	"crud_echo/shared/response"
	"github.com/labstack/echo/v4"
	"net/http"
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
