package usecase

import "github.com/koyarskaya/laba-11/internal/auth/api"

type Provider interface {
	CheckUser(api.User) (api.User, error)
	CreateUser(api.User) error
	SelectUser(string) (api.User, error)
}
