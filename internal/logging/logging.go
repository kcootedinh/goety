package logging

import (
	"log/slog"
	"os"
)

var baseLogger *slog.Logger

func init() {
	baseLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With(ParamApp, AppName)
}

func New(verbose bool) *slog.Logger {
	if baseLogger == nil {
		// should not occur
		baseLogger = slog.Default()
	}

	if verbose {
		baseLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	}

	return baseLogger
}
