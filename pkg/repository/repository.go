package repository

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
func NewRepository() *Repository {
	return &Repository{}
}
