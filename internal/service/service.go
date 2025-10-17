package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// проверяем, состоит ли полученная строка только из символов кода Морзе
func isMorse(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false // строка пустая и не содержит каких-либо символов
	}
	for _, r := range s {
		switch r {
		case '.', '-', ' ', '/':
			continue
		default:
			return false
		}
	}
	return true
}

// Convert автоматически определяет тип даных и конвертирует их
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("input string is empty")
	}

	var result string

	// автоматически определяем входные данные и конвертируем в противоположный формат
	if isMorse(input) {
		result = morse.ToText(input)
	} else {
		result = morse.ToMorse(input)
	}

	// проверяем на наличие во входной строке неподдерживаемых символов
	if result == "" && input != "" {
		return "", errors.New("contains unsupported symbols")
	}

	return result, nil
}
