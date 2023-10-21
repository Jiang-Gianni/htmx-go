package server

import (
	"database/sql"
	"io/fs"
	"log"
	"net/http"

	"github.com/Jiang-Gianni/htmx-go/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

// Common variables
var (
	err          error
	query        *db.Queries
	routerGroups = []func(r chi.Router){
		indexGroup,
		firstGroup,
	}
)

func init() {
	query = db.New(InitSqlite())
}

func InitSqlite() *sql.DB {
	sqlite, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	return sqlite
}

// Parse html templates, register http routings and serve on port 3000
func RegisterAndRun(assetsFs fs.FS) {

	// Chi Router
	r := Router(assetsFs)

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
		// w.Header().Set("Cache-Control", "max-age=2592000")
		w.Header().Set("Cache-Control", "no-store")
		fs := http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFs)))
		fs.ServeHTTP(w, r)
	})

	// Mapping
	for _, group := range routerGroups {
		router.Group(group)
	}

	return router
}
