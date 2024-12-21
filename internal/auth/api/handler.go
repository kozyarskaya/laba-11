package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrEmailAlreadyTaken = errors.New("email уже занят")

//middleware - функция, которая при обработке запросов может выполнять доп действия до или после вызова обработчика

// проверяет наличие и валидность JWT в заголовке Authorization каждого запроса
// next - следдующий обработчик
// возвращает новую функцию, соответствующую типу echo.HandlerFunc, которая будет использоваться как middleware

// signUp - обработчик для регистрации пользователя
func (srv *Server) signUp(c echo.Context) error {
	fmt.Println("user signUp from api")
	var user User
	if err := c.Bind(&user); err != nil {
		fmt.Println("ppp")
		return c.JSON(http.StatusBadRequest, Response{
			Message: "некорекктное считывание данных",
		})
	}
	fmt.Println("user signUp from api")
	// Вызываем бизнес-логику для регистрации пользователя
	token, err := srv.uc.SignUp(user)
	if err != nil {
		fmt.Println(err.Error(), errors.Is(err, ErrEmailAlreadyTaken))
		if errors.Is(err, ErrEmailAlreadyTaken) { // Проверяем конкретную ошибку
			fmt.Println("rrr")
			return c.JSON(http.StatusConflict, Response{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, Response{
			Message: "ошибка сервера",
		})
	}
	fmt.Println("signUp sucs")

	return c.JSON(http.StatusCreated, Response{
		Message: token,
	})
}

// Обработчик для аутентификации пользователя
func (srv *Server) signIn(c echo.Context) error {
	var credentials Credentials
	if err := c.Bind(&credentials); err != nil {
		fmt.Println("Invalid input")
		return c.JSON(http.StatusBadRequest, Response{
			Message: " Ошибка передачи параметров",
		})
	}
	token, err := srv.uc.SignIn(credentials)
	if err != nil {
		fmt.Println("Authentication failed")
		return c.JSON(http.StatusUnauthorized, Response{
			Message: " Ошибка сервера",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Message: token,
	})
}

func (srv *Server) hello(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Message: "Hello!",
	})
}
