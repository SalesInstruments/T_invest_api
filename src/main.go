package main

import (
	g "T_invest_api/internal/globals"
	router "T_invest_api/internal/http-server/routes"
	"net/http"

	"golang.org/x/exp/slog"
)

func main() {

	g.Log.Info(
		"start T_invest_api",
		slog.String("env", g.Cfg.Env),
		slog.String("adress", g.Cfg.HTTPServer.Address),
	)

	srv := &http.Server{
		Addr:         g.Cfg.HTTPServer.Address,
		Handler:      router.New(),
		ReadTimeout:  g.Cfg.HTTPServer.Timeout,
		WriteTimeout: g.Cfg.HTTPServer.Timeout,
		IdleTimeout:  g.Cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		g.Log.Error("failed to start server")
	}

	g.Log.Error("server stoped")

}
