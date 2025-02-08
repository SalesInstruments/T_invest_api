package main

import (
	"T_invest_api/internal/config"
	mvLoger "T_invest_api/internal/http-server/middleware/logger"
	"T_invest_api/internal/logger"
	"T_invest_api/internal/storage"

	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	cfg = config.MustLoad()
	log = logger.SetupLogger(cfg.Env)
)

func main() {

	log.Info("start T_invest_api", slog.String("env", cfg.Env))
	log.Debug("debag messages are enable")

	storage, err := storage.New()
	if err != nil {
		log.Error("failed to init storage", logger.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mvLoger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	_ = storage
}
