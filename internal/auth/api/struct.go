package api

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id             int    `json:"-"`
	Name           string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
}

//LoginRequest
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}
