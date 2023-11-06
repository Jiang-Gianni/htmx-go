package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *Api) MountHandlers() *chi.Mux {
	router := chi.NewRouter()

	// Middlewares
	// router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Serving compressed static files
	router.Get("/assets/*", a.getAssets())

	// Pprof
	router.Mount("/debug", middleware.Profiler())

	// Index
	router.Get("/", a.getIndex())
	router.Get("/first", a.getFirst())

	return router
}
