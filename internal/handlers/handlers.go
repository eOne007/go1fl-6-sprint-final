package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// MainHandler возвращает HTML из файла
func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler загружает файл, парсит форму, конвертирует текст, создает файл и записывает в него результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// парсим multipart-форму (макс размер - 10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "HTML parsing error", http.StatusInternalServerError)
		return
	}
	// получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// читаем содержимое полученного файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "read file error", http.StatusInternalServerError)
		return
	}

	// передаем содержимое файла в функцию пакета Service (service.Convert)
	convertFile := string(data)
	result, err := service.Convert(convertFile)
	if err != nil {
		http.Error(w, "convertation error", http.StatusInternalServerError)
		return
	}

	// создаем расширение и наименование файла
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt" // если расширение не указано
	}
	filename := time.Now().UTC().Format("2006-01-02_15-04-05.000000") + ext

	// создаем локальный файл
	localFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "file creating error", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	// записываем результат в созданный файл
	_, err = localFile.WriteString(result)
	if err != nil {
		http.Error(w, "writing to file failed", http.StatusInternalServerError)
		return
	}
	// возвращаем результат конвертации данных
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
