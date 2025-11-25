package core

import "net/http"

// Logger интерфейс для логгеров
type Logger interface {
	Info(msg string, attrs ...any)
	Warn(msg string, attrs ...any)
	Error(msg string, attrs ...any)
	Debug(msg string, attrs ...any)
}

// HTTPClient интерфейс для HTTP-клиентов
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
