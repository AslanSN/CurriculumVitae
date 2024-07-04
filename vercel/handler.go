package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/AslanSN/CurriculumVitae/views"
	"github.com/a-h/templ"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/static/") {
		// Servir el archivo estático
		log.Printf("Serving static file: %s", r.URL.Path)

		http.FileServer(http.Dir("./assets")).ServeHTTP(w, r)
		return
	}
	h := templ.Handler(views.Index())
	h.ServeHTTP(w, r)
}
