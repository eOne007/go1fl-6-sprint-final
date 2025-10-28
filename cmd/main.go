package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// создаём логгер и сервер
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	srv := server.NewServer(logger)

	// Запускаем HTTP-сервер
	err := srv.Start()
	if err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}

	logger.Println("Сервер остановлен")
}
