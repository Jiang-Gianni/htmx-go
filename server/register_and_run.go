package server

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Common variables
var (
	err          error
	routerGroups = []func(r chi.Router){
		indexGroup,
		firstGroup,
	}
)

// Parse html templates, register http routings and serve on port 3000
func RegisterAndRun(assetsFs fs.FS) {

	// Chi Router
	r := Router(assetsFs)

	log.Println("listening on port 3000")
	if err = http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}

func Router(assetsFs fs.FS) *chi.Mux {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Serving static files
	router.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFs)))
		fs.ServeHTTP(w, r)
	})

	// Mapping
	for _, group := range routerGroups {
		router.Group(group)
	}

	return router
}
