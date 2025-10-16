package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// создаем структуру сервера с полями для логгера и http-сервера
type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// NewServer создает и настраивает новый сервер
func NewServer(logger *log.Logger) *Server {

	// создаем новый роутер
	mux := http.NewServeMux()

	// регистрируем в роутере хендлеры из пакета handlers
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// создаем экземпляр структуры с перечислением полей согласно ТЗ

	httpServer := &http.Server{
		Addr:         ":8080",          // используем порт 8080
		Handler:      mux,              // передаем наш http-роутер
		ErrorLog:     logger,           // передлаем наш логгер
		ReadTimeout:  5 * time.Second,  // таймаут для чтения - 5 секунд
		WriteTimeout: 10 * time.Second, // таймаут для записи - 10 секунд
		IdleTimeout:  15 * time.Second, // таймаут ожидания следующего запроса - 15 секунд
	}

	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}

}
