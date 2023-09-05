package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
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
	return &Repository{}
}
