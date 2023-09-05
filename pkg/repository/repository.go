package repository

import (
	todo "todo_app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

// собирает все сервисы в одном месте
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// объявляем конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
