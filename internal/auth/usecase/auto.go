package usecase

// реализуем бизнес-логику, ей вызывают обрабтчики запросов, а она передает задачу бд
import (
	"fmt"

	"github.com/kozyarskaya/laba-11/internal/auth/api"
	"github.com/kozyarskaya/laba-11/utils"
)

func (u *Usecase) SignUp(user api.User) (string, error) {
	// Проверяем, существует ли уже пользователь с таким email
	existingUser, err := u.p.CheckUser(user)
	if err != nil {
		fmt.Println(err)
		return "", err // Обработка ошибки при проверке пользователя
	}
	fmt.Println("checked user")
	if existingUser.Email != "" {
		fmt.Println("email")
		return "", api.ErrEmailAlreadyTaken // Возвращаем конкретное сообщение об ошибке
	}
	// Хешируем пароль перед сохранением
	hashedPassword, err := utils.HashPassword(user.HashedPassword) // Измените на user.Password
	if err != nil {
		fmt.Println("pas eror")
		return "", err
	}
	newUser := user
	newUser.HashedPassword = hashedPassword

	// Сохраняем нового пользователя в базу данных
	err = u.p.CreateUser(newUser)
	if err != nil {
		fmt.Println("error of create")
		return "", err
	}
	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return "", err // Возвращаем ошибку при генерации токена
	}
	fmt.Println("SignUp ssuccsesf")
	return token, nil
}

// SignIn аутентифицирует пользователя и возвращает JWT токен
func (u *Usecase) SignIn(credentials api.Credentials) (string, error) {
	var user api.User
	// Получаем пользователя по email
	user, err := u.p.SelectUser(credentials.Email)
	if err != nil {
		fmt.Println("user not found")
		return "", err
	}
	// Сравниваем хешированный пароль с введенным паролем
	if err := utils.ComparePasswords(user.HashedPassword, credentials.Password); err != nil {
		fmt.Println("password not compared", utils.ComparePasswords(user.HashedPassword, credentials.Password))
		return "", err
	}
	fmt.Println("password is correct")
	// Генерация JWT токена после успешной аутентификации
	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return "", err // Возвращаем ошибку при генерации токена
	}
	return token, nil // Возвращаем сгенерированный токен
}
