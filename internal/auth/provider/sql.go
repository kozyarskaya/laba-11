package provider

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kozyarskaya/laba-11/internal/auth/api"
)

func (p *Provider) CheckUser(u api.User) (api.User, error) {
	var user api.User
	err := p.conn.QueryRow(`SELECT id, email, name FROM users WHERE email = $1`, u.Email).Scan(&user.Id, &user.Email, &user.Name)

	fmt.Println("check user")

	if err != nil {
		if err == sql.ErrNoRows {
			// Если пользователь не найден, возвращаем пустого пользователя и nil как ошибку
			return api.User{}, nil
		}
		// Возвращаем ошибку при проблеме с запросом
		return api.User{}, fmt.Errorf("ошибка при проверке пользователя: %w", err)
	}
	fmt.Println("checked from check")
	// Если пользователь найден, возвращаем его и ошибку о том, что почта занята
	return user, nil // Возвращаем найденного пользователя без ошибки
}

// Определение ошибки для занятости email
var ErrEmailAlreadyTaken = errors.New("email уже занят")

func (p *Provider) CreateUser(u api.User) error {
	_, err := p.conn.Exec("INSERT INTO users (name, email, hashedPassword) VALUES ($1, $2, $3)", u.Name, u.Email, u.HashedPassword)
	fmt.Println("create user")
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) SelectUser(emai string) (api.User, error) {
	var user api.User
	err := p.conn.QueryRow("SELECT id, name, email, hashedPassword FROM users WHERE email = $1", emai).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.User{}, errors.New("invalid credentials") // Пользователь не найден
		}
		return api.User{}, err // Возвращаем ошибку при проблеме с запросом
	}
	return user, nil
}
