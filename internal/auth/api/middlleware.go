package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	//контекст содержит информацию о текущем HTTP-запросе и позволяет взаимодействовать с ним
	fmt.Println("ppp")
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		token = strings.TrimPrefix(token, "Bearer ") //удаление префикса
		claims := &Claims{}                          // Claims - это структура, которая содержит данные о пользователе
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !tkn.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		c.Set("userId", claims.UserId) // Сохраняем ID пользователя в контексте
		return next(c)
	}
}
func TokenValidationHandler(c echo.Context) error {
	token := c.QueryParam("token")
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing token"})
	}

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return "your_secret_key", nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "valid token"})
}
