package logger

import (
	"context"
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	log.Info("url-shortener", slog.String("env", env))

	if log.Handler().Enabled(context.Background(), slog.LevelDebug) {
		log.Debug("debug level", slog.String("env", env))
	}

	return log
}
