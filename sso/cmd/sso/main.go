package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	// TODO инициализация объекта конфига
	cfg := config.MustLoad()

	fmt.Println(*cfg)
	// TODO инициалзаация логгера
	// TODO инициализация приложения (app)
	// TODO  запустит gRPC сервер
}
