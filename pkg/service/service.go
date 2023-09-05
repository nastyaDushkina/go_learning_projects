package service

import (
	todo "todo_app"
	"todo_app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

// собирает все сервисы в одном месте
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// объявляем конструктор
// (внедрение зависимостей)
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
