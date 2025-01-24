package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/NikCool98/url-short/internal/config"
	"github.com/NikCool98/url-short/internal/config/lib/logger/sl"
	"github.com/NikCool98/url-short/internal/config/storage"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug message are enabled")
	storage, err := storage.ConnectDB(cfg)
	if err != nil {
		log.Error("Failed to connect DB", sl.Err(err))
		return
	}
	defer storage.DB.Close()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
