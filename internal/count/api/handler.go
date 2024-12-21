package api

import (
	"errors"
	"net/http"

	"github.com/kozyarskaya/laba-11/pkg/vars"
	"github.com/labstack/echo/v4"
)

type CountInput struct {
	Val float32 `json:"val"` // Используем float32 для автоматической проверки числового значения
}

func (srv *Server) GetCount(e echo.Context) error {
	msg, err := srv.uc.FetchCount()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) PostCount(c echo.Context) error {
	input := CountInput{}
	// Привязка входных данных и проверка на ошибки
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "Неправильный формат JSON")
	}
	err := srv.uc.SetCount(input.Val)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return c.String(http.StatusConflict, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, "OK")
}

func (srv *Server) PutCount(c echo.Context) error {
	input := CountInput{}
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "Неправильный формат JSON")
	}
	err := srv.uc.ChangeCount(input.Val)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return c.String(http.StatusConflict, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, "OK")
}
