package server

import (
	"net/http"

	"github.com/Jiang-Gianni/htmx-go/views"
	"github.com/go-chi/chi/v5"
)

var (
	indexPage = &views.Page{
		Title:       "Title HTMX and Go",
		Description: "Index Page",
	}
	indexGroup = func(r chi.Router) {
		r.Get("/", getIndex)
	}
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	views.WriteIndexPage(w, indexPage, views.IndexBody())
}
