package todo

// можно хранить структуры на самом верхнем уровне проекта

type User struct {
	Id int `json:"-"`
	// валидация наличия полей в теле запросаб реализация фреймворка gin
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
