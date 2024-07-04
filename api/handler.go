package handler

import (
	"net/http"
	"strings"

	"github.com/AslanSN/CurriculumVitae/views"
	"github.com/a-h/templ"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/static/") {
		// Servir el archivo estático
		http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
		return
	}
	h := templ.Handler(views.Index())
	h.ServeHTTP(w, r)
}
