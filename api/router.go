package api

import (
	"net/http"

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
	router.Use(a.AllowedMiddeware)

	// Serving compressed static files
	router.Get("/assets/*", a.getAssets())

	// Pprof
	router.Mount("/debug", middleware.Profiler())

	// Index
	router.Get("/", a.getIndex())
	router.Get("/first", a.getFirst())

	return router
}

func (a *Api) AllowedMiddeware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !a.Allowed(*r) {
				w.WriteHeader(http.StatusForbidden)
				a.log.Warn.Println("Invalid resource access")
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}

// Anti CSRF
func (a *Api) Allowed(r http.Request) bool {
	site := r.Header.Get("sec-fetch-site")
	mode := r.Header.Get("sec-fetch-mode")
	// Same site or direct url or not supported by browser
	if site == "" || site == "none" || site == "same-site" || site == "same-origin" {
		return true
	}
	// Cross site
	if mode == "navigate" && r.Method == http.MethodGet {
		return true
	}
	return false
}
