package main

import (
	"T_invest_api/internal/config"
	router "T_invest_api/internal/http-server/routes"
	"T_invest_api/internal/logger"
	"net/http"

	"golang.org/x/exp/slog"
)

var (
	cfg = config.MustLoad()
	log = logger.SetupLogger(cfg.Env)
)

func main() {

	log.Info(
		"start T_invest_api",
		slog.String("env", cfg.Env),
		slog.String("adress", cfg.HTTPServer.Address),
	)

	// storage, err := storage.New()
	// if err != nil {
	// 	log.Error("failed to init storage", logger.Err(err))
	// 	os.Exit(1)
	// }
	// _ = storage

	// router := chi.NewRouter()

	// router.Use(middleware.RequestID)
	// router.Use(mwLoger.New(log))
	// router.Use(middleware.Recoverer)
	// router.Use(middleware.URLFormat)

	// router.Post("/", bonds.New())

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router.New(),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stoped")

}
