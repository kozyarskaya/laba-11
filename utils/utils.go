package utils

import (
	"fmt"
	"time"

	"github.com/kozyarskaya/laba-11/internal/auth/api"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key") // ваш секретный ключ

// HashPassword хеширует пароль с помощью bcrypt
func HashPassword(password string) (string, error) {
	fmt.Println(password, "dddd")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("password for hash", password, string(bytes))
	return string(bytes), err
}

// GenerateToken создает новый JWT-токен
func GenerateToken(userID int) (string, error) {
	claims := &api.Claims{
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Токен истекает через 24 часа
			Issuer:    "your_app_name",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ComparePasswords сравнивает хешированный пароль с введенным паролем
func ComparePasswords(hashedPassword string, password string) error {
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
