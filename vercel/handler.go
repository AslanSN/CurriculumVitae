package handler

import (
	"net/http"

	"github.com/AslanSN/CurriculumVitae/app/views"
	"github.com/a-h/templ"
)

// Handler is the Vercel (@vercel/go) entrypoint. Static assets under /assets are
// served by Vercel's CDN (a @vercel/static build + the "filesystem" route in
// vercel.json), so this only needs to render the page for non-asset paths.
func Handler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.Index()).ServeHTTP(w, r)
}
