package main

import (
	"T_invest_api/internal/config"
	"T_invest_api/internal/http-server/handlers/url/bonds"
	mvLoger "T_invest_api/internal/http-server/middleware/logger"
	"T_invest_api/internal/logger"
	"net/http"

	"log/slog"

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

	// storage, err := storage.New()
	// if err != nil {
	// 	log.Error("failed to init storage", logger.Err(err))
	// 	os.Exit(1)
	// }
	// _ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mvLoger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/", bonds.New())

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stoped")

}
