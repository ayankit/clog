package clog

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

var log *slog.Logger

func Debug(msg string, args ...any) { log.Debug(msg, args...) }
func Info(msg string, args ...any)  { log.Info(msg, args...) }
func Warn(msg string, args ...any)  { log.Warn(msg, args...) }
func Error(msg string, args ...any) { log.Error(msg, args...) }

func Debugf(format string, args ...any) { log.Debug(fmt.Sprintf(format, args...)) }
func Infof(format string, args ...any)  { log.Info(fmt.Sprintf(format, args...)) }
func Warnf(format string, args ...any)  { log.Warn(fmt.Sprintf(format, args...)) }
func Errorf(format string, args ...any) { log.Error(fmt.Sprintf(format, args...)) }

func With(args ...any) *slog.Logger { return log.With(args...) }

func Fatal(msg string, args ...any) {
	log.Error(msg, args...)
	os.Exit(1)
}
func Fatalf(format string, args ...any) {
	log.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Init(level Level, logFilePath ...string) error {
	var handler slog.Handler

	opts := Options{
		Level: level.toSlogLevel(),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Value.Kind() == slog.KindAny {
				if _, ok := a.Value.Any().(error); ok {
					return Attr(9, a)
				}
			}
			return a
		},
		TimeFormat: "2006-01-02 15:04:05.000",
	}

	if len(logFilePath) == 1 {
		if err := os.MkdirAll(filepath.Dir(logFilePath[0]), 0755); err != nil {
			return err
		}

		file, err := os.OpenFile(logFilePath[0], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		writer := io.MultiWriter(os.Stdout, file)
		handler = NewHandler(writer, &opts)
	} else {
		handler = NewHandler(os.Stdout, &opts)
	}

	log = slog.New(handler)
	slog.SetDefault(log)
	return nil
}
