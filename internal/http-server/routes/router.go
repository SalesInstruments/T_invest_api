package router

import (
	"T_invest_api/internal/config"
	"T_invest_api/internal/http-server/handlers/url/instrument"
	marketdata "T_invest_api/internal/http-server/handlers/url/marketData"
	mwLoger "T_invest_api/internal/http-server/middleware/logger"
	"T_invest_api/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	cfg = config.MustLoad()
	log = logger.SetupLogger(cfg.Env)
)

type Router struct {
	*chi.Mux
}

func New() *Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLoger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	r := &Router{router}

	r.appendRout()

	return r
}

func (r *Router) appendRout() {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/shareBy", instrument.New())
			r.Get("/getCandles", marketdata.New())
		})

		// r.Route("/v1", func(r chi.Router) {

		// })
	})
}
