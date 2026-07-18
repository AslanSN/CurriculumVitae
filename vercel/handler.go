package handler

import (
	"net/http"

	"github.com/AslanSN/CurriculumVitae/app/views"
	"github.com/AslanSN/CurriculumVitae/i18n"
)

// Handler is the Vercel (@vercel/go) entrypoint. Static assets under /assets are
// served by Vercel's CDN (a @vercel/static build + the "filesystem" route in
// vercel.json), so this only needs to serve robots/sitemap, resolve the locale,
// and render the page.
func Handler(w http.ResponseWriter, r *http.Request) {
	// SEO files are locale-independent and must bypass the locale redirect so a
	// crawler's Accept-Language never bounces /sitemap.xml to /es.
	switch r.URL.Path {
	case "/robots.txt":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(i18n.RobotsTXT()))
		return
	case "/sitemap.xml":
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		_, _ = w.Write([]byte(i18n.SitemapXML()))
		return
	}

	locale, redirect := i18n.Resolve(r)
	if redirect != "" && redirect != r.URL.Path {
		http.Redirect(w, r, redirect, http.StatusFound)
		return
	}

	// Remember the resolved locale so root visits stay consistent with the
	// visitor's language (auto-detected once, or pinned by the switcher).
	i18n.SetCookie(w, locale)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx := i18n.WithLocale(r.Context(), locale)
	_ = views.Index().Render(ctx, w)
}
