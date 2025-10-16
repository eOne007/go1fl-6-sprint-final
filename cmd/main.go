package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// создаём логгер и сервер
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)
	srv := server.NewServer(logger)

	// логируем запуск
	logger.Println("Сервер запускается на :8080...")

	// Запускаем HTTP-сервер
	err := srv.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}

	logger.Println("Сервер остановлен")
}
