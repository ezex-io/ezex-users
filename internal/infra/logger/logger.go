package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewSlog(handler slog.Handler) *Logger {
	if handler == nil {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}
	return &Logger{
		Logger: slog.New(handler),
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}
