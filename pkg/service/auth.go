package service

import (
	"crypto/sha1"
	"fmt"
	todo "todo_app"
	"todo_app/pkg/repository"
)

const salt = "hgsgjqjqnncbx1d5d4e6ds1"

// реализация интерфейса
type AuthService struct {
	repo repository.Authorization
}

// конструктор
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	// хэшируем пароль
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
