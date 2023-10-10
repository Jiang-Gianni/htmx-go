package server

import (
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
func RegisterAndRun() {

	// Chi Router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(htmlContentType)

	// Serving static files
	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
		fs.ServeHTTP(w, r)
	})

	// Mapping
	for _, group := range routerGroups {
		r.Group(group)
	}

	println("http://localhost:3000/")
	if err = http.ListenAndServe(":3000", r); err != nil {
		println(err.Error())
	}

}
