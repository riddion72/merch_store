package logger

import (
	"log/slog"
	"os"
	"sync"
)

type LogWrapper struct {
	logger *slog.Logger
	sync.Once
}

var log *LogWrapper

func MustInitlogger(level string) {
	log = &LogWrapper{}
	log.Do(func() {
		switch level {
		case "local":
			log.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case "dev":
			log.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		case "prod":
			log.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		}
	})
}

func Debug(msg string, fields ...slog.Attr) {
	var args []any
	for i := range fields {
		args = append(args, slog.Attr{Key: fields[i].Key, Value: slog.StringValue(fields[i].Value.String())})
	}
	l := log.logger.With(args...)
	l.Debug(msg)
}

func Info(msg string, fields ...slog.Attr) {
	var args []any
	for i := range fields {
		args = append(args, slog.Attr{Key: fields[i].Key, Value: slog.StringValue(fields[i].Value.String())})
	}
	l := log.logger.With(args...)
	l.Info(msg)
}

func Warn(msg string, fields ...slog.Attr) {
	var args []any
	for i := range fields {
		args = append(args, slog.Attr{Key: fields[i].Key, Value: slog.StringValue(fields[i].Value.String())})
	}
	l := log.logger.With(args...)
	l.Warn(msg)
}

func Error(msg string, fields ...slog.Attr) {
	var args []any
	for i := range fields {
		args = append(args, slog.Attr{Key: fields[i].Key, Value: slog.StringValue(fields[i].Value.String())})
	}
	l := log.logger.With(args...)
	l.Error(msg)
}
