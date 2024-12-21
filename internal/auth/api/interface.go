package api

type Usecase interface {
	SignUp(User) (string, error)
	SignIn(Credentials) (string, error) //Возвращает JWT токен,
}
