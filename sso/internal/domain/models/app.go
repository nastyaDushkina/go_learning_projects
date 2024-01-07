package models

type App struct {
	ID     int
	Name   string
	Secret string // секрет нужен для подписи токенпа и для дальнейшей валидации токена на стороне клиента
}
