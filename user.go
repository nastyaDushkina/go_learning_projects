package todo

// можно хранить структуры на самом верхнем уровне проекта

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
