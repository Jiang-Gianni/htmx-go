package vercel

import (
	"net/http"
	"os"

	"github.com/Jiang-Gianni/htmx-go/server"
	"github.com/go-chi/chi/v5"
)

var (
	router *chi.Mux
)

func init() {
	router = server.Router(
		os.DirFS("assets"),
	)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
