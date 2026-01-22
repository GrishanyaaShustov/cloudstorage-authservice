// Package logger provides application-wide logging setup and helpers.
package logger

import (
	"log/slog"
	"os"

	"github.com/GrishanyaaShustov/cloudstorage-authservice/pkg/logger/handlers/slogpretty"
)

// SetupLogger initializes and returns a slog.Logger
// configured according to the provided environment.
func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = setupPrettySlog()
	case envProd:
		log = setupPrettySlog()
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
