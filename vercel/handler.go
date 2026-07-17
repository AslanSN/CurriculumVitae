package handler

import (
	"net/http"
	"strings"

	cv "github.com/AslanSN/CurriculumVitae"
	"github.com/AslanSN/CurriculumVitae/app/views"
	"github.com/a-h/templ"
)

// Handler is the Vercel (@vercel/go) entrypoint. It serves the embedded static
// assets under /assets (the FS keeps the "assets/" prefix, so the request path
// maps 1:1 with no rewriting) and renders the CV page for every other path.
func Handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/assets/") {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		http.FileServerFS(cv.Assets).ServeHTTP(w, r)
		return
	}
	templ.Handler(views.Index()).ServeHTTP(w, r)
}
