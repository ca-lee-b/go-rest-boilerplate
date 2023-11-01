package log

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return logger
}
