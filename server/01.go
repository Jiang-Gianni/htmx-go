package server

import (
	"net/http"

	"github.com/Jiang-Gianni/htmx-go/views"
	"github.com/go-chi/chi/v5"
)

var (
	firstPage = &views.Page{
		Title:       "First",
		Description: "Index Page",
	}
	firstGroup = func(r chi.Router) {
		r.Get("/first", getFirst)
	}
)

func getFirst(w http.ResponseWriter, r *http.Request) {
	views.WriteIndexPage(w, firstPage, views.First())
}
