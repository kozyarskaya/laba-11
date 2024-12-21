package api

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/kozyarskaya/laba-11/pkg/vars"
	"github.com/labstack/echo/v4"
)

// GetQuery возвращает приветствие случайному пользователю
func (srv *Server) GetQuery(e echo.Context) error {
	msg, err := srv.uc.FetchQuery()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

// PostQuery Помещает нового пользователя в БД
func (srv *Server) PostQuery(c echo.Context) error {
	nameInput := c.QueryParam("name") // Получаем Query-параметр
	if nameInput == "" {
		return c.String(http.StatusBadRequest, "Missing 'name' query parameter")
	}
	//проверка на корректность имени
	re := regexp.MustCompile(`[a-zA-Zа-яА-Я]`)
	if !re.MatchString(nameInput) {
		return c.String(http.StatusBadRequest, "empty string")
	}
	if len([]rune(nameInput)) > srv.maxSize {
		return c.String(http.StatusBadRequest, "name too large")
	}
	err := srv.uc.SetQuery(nameInput)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return c.String(http.StatusConflict, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, "OK")
}

func (srv *Server) PutQuery(c echo.Context) error {
	nameInput := c.QueryParam("name") // Получаем Query-параметр
	if nameInput == "" {
		return c.String(http.StatusBadRequest, "Missing 'name' query parameter")
	}
	//проверка на корректность имени
	re := regexp.MustCompile(`[a-zA-Zа-яА-Я]`)
	if !re.MatchString(nameInput) {
		return c.String(http.StatusBadRequest, "empty string")
	}
	if len([]rune(nameInput)) > srv.maxSize {
		return c.String(http.StatusBadRequest, "name too large")
	}
	err := srv.uc.ChangeQuery(nameInput)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return c.String(http.StatusConflict, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusCreated, "OK")
}
