package logger

import (
	"log/slog"
	"sync"
)

var logger_instance *logger

type ILogger interface {
	Info(string)
	Debug(string)
	Error(string, ...any)
}

type logger struct {
}

func NewLogger() *logger {
	var once sync.Once
	once.Do(func() {
		logger_instance = &logger{}
	})
	return logger_instance
}

func (l *logger) Info(message string) {
	slog.Info(message)
}

func (l *logger) Debug(message string) {
	slog.Debug(message)
}

func (l *logger) Error(message string, args ...any) {
	slog.Error(message, args)
}
