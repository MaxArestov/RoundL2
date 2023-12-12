package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(h http.Handler) *Logger {
	return &Logger{h}
}

func (l *Logger) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(writer, request)
	// Метод, урла и сколько прошло времени с запуска логгера.
	log.Printf("%s; %s; %v", request.Method, request.URL.Path, time.Since(start))
}
