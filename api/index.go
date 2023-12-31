package api

import (
	"net/http"

	"github.com/Jiang-Gianni/htmx-go/views"
)

func (a *Api) getIndex() http.HandlerFunc {
	indexPage := &views.Page{
		Title:       "Title HTMX and Go",
		Description: "Index Page",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		views.WriteIndexPage(w, indexPage, views.IndexBody())
	}
}
