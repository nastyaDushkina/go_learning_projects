package main

import (
	"log"
	todo "todo_app" //todo package
	"todo_app/pkg/handler"
)

func main() {
	// для запуска сервера у эндпоинтов должен быть хотя бы один обработчик
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
