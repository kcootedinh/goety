package logging

import (
	"log/slog"
	"os"
)

var baseLogger *slog.Logger

func init() {
	baseLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})).With(ParamApp, AppName)
}

func Logger() *slog.Logger {
	if baseLogger == nil {
		// should not occur
		return slog.Default()
	}

	return baseLogger
}
