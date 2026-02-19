package clog_test

import (
	"errors"
	"time"

	"github.com/ayankit/clog"
)

func Example() {
	clog.Init(clog.LevelDebug)

	clog.Info("Starting server", "addr", ":8080", "env", "production")
	clog.Debug("Connected to DB", "db", "myapp", "host", "localhost:5432")
	clog.Warn("Slow request", "method", "GET", "path", "/users", "duration", 497*time.Millisecond)
	clog.Error("DB connection lost", clog.Err(errors.New("connection reset")), "db", "myapp")
}

// Create a new logger that writes to given log file
func Example_withLogToFile() {
	clog.Init(clog.LevelDebug, "/path/to/log/file")

	clog.Info("Starting server", "addr", ":8080", "env", "production")
	clog.Debug("Connected to DB", "db", "myapp", "host", "localhost:5432")
	clog.Warn("Slow request", "method", "GET", "path", "/users", "duration", 497*time.Millisecond)
	clog.Error("DB connection lost", clog.Err(errors.New("connection reset")), "db", "myapp")
}
