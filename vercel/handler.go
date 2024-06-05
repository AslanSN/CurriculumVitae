package handler

import (
	"net/http"

	"github.com/AslanSN/CV/views"
	"github.com/a-h/templ"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	h := templ.Handler(views.Index())
	h.ServeHTTP(w, r)
}
