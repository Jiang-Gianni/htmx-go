package api

import (
	"net/http"

	"github.com/Jiang-Gianni/htmx-go/views"
)

func (a *Api) getFirst() http.HandlerFunc {
	firstPage := &views.Page{
		Title:       "First",
		Description: "Index Page",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		views.WriteIndexPage(w, firstPage, views.First())
	}
}
