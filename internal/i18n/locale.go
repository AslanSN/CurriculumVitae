// Package i18n provides locale detection, per-request locale plumbing through
// context, and a type-safe UI string catalog for the trilingual (EN/ES/FR) CV.
package i18n

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Locale is a supported UI language, identified by its ISO 639-1 code.
type Locale string

const (
	EN Locale = "en"
	ES Locale = "es"
	FR Locale = "fr"
)

// Default is served at the root path ("/") and used as the fallback whenever a
// locale can't be resolved. English is the canonical language for SEO.
const Default = EN

// Supported lists locales in the order they appear in the language switcher.
var Supported = []Locale{EN, ES, FR}

// cookieName stores the visitor's last explicit language choice so the root
// path keeps honoring it instead of re-detecting on every visit.
const cookieName = "lang"

// Valid reports whether l is one of the supported locales.
func (l Locale) Valid() bool {
	for _, s := range Supported {
		if s == l {
			return true
		}
	}
	return false
}

// Lang is the value for the <html lang> attribute.
func (l Locale) Lang() string { return string(l) }

// Native is the endonym shown in the language switcher.
func (l Locale) Native() string {
	switch l {
	case ES:
		return "Español"
	case FR:
		return "Français"
	default:
		return "English"
	}
}

// Short is the uppercase code for compact UI (EN/ES/FR).
func (l Locale) Short() string { return strings.ToUpper(string(l)) }

// Path is the canonical URL for the locale — "/" for the default (English),
// "/es" and "/fr" for the others.
func (l Locale) Path() string {
	if l == Default {
		return "/"
	}
	return "/" + string(l)
}

// SiteURL is the canonical production origin, used to build absolute SEO URLs.
const SiteURL = "https://aslansn-cv.vercel.app"

// CanonicalURL is the absolute canonical URL for the locale (for <link rel>
// canonical / alternate hreflang tags).
func (l Locale) CanonicalURL() string { return SiteURL + l.Path() }

// OpenGraph is the og:locale value (language_TERRITORY) for the locale.
func (l Locale) OpenGraph() string {
	switch l {
	case ES:
		return "es_ES"
	case FR:
		return "fr_FR"
	default:
		return "en_US"
	}
}

// SwitchPath is the href the language switcher points at to move to l. The
// default locale carries an explicit ?lang override so the choice pins the
// cookie and root auto-detection won't bounce it back; the others use their
// clean canonical path.
func SwitchPath(l Locale) string {
	if l == Default {
		return "/?lang=" + string(l)
	}
	return "/" + string(l)
}

// --- context plumbing -------------------------------------------------------

type ctxKey struct{}

// WithLocale returns a copy of ctx carrying the locale. templ components read
// it back with FromContext during rendering.
func WithLocale(ctx context.Context, l Locale) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}

// FromContext returns the locale stored in ctx, or Default if absent/invalid.
func FromContext(ctx context.Context) Locale {
	if l, ok := ctx.Value(ctxKey{}).(Locale); ok && l.Valid() {
		return l
	}
	return Default
}

// --- request resolution -----------------------------------------------------

// Resolve determines the locale for a request and, for the root path, whether
// the visitor should be redirected to the canonical localized URL.
//
// Precedence: explicit path prefix (/es, /fr) > ?lang= override > lang cookie >
// Accept-Language. The returned redirect is non-empty only when a non-default
// locale is resolved for a request that isn't already on that locale's path, so
// /es and /fr always get their own canonical URL (better for SEO and sharing).
func Resolve(r *http.Request) (locale Locale, redirect string) {
	path := r.URL.Path

	// 1. An explicit locale in the path always wins — no redirect needed.
	switch {
	case path == "/es" || strings.HasPrefix(path, "/es/"):
		return ES, ""
	case path == "/fr" || strings.HasPrefix(path, "/fr/"):
		return FR, ""
	}

	// Any other path is treated as the English root. Decide whether a
	// non-default locale should take over (and redirect to its own URL).
	pick := func(l Locale) (Locale, string) {
		if l == Default {
			return EN, ""
		}
		return l, l.Path()
	}

	// 2. Explicit override from the language switcher pins the choice.
	if q := Locale(strings.ToLower(r.URL.Query().Get("lang"))); q.Valid() {
		return pick(q)
	}

	// 3. A previously chosen language, remembered in a cookie.
	if c, err := r.Cookie(cookieName); err == nil {
		if cl := Locale(c.Value); cl.Valid() {
			return pick(cl)
		}
	}

	// 4. Auto-detection from the browser's Accept-Language header.
	return pick(detectAcceptLanguage(r.Header.Get("Accept-Language")))
}

// SetCookie persists the resolved locale so subsequent root visits stay
// consistent with the visitor's choice.
func SetCookie(w http.ResponseWriter, l Locale) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    string(l),
		Path:     "/",
		MaxAge:   int((365 * 24 * time.Hour).Seconds()),
		HttpOnly: false, // readable by client so a future JS switcher can reuse it
		SameSite: http.SameSiteLaxMode,
	})
}

// detectAcceptLanguage returns the first supported locale referenced in an
// Accept-Language header, honoring the quality-value ordering, or Default.
func detectAcceptLanguage(header string) Locale {
	if header == "" {
		return Default
	}

	type ranked struct {
		locale Locale
		q      float64
	}
	var best *ranked

	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		tag := part
		q := 1.0
		if i := strings.Index(part, ";"); i != -1 {
			tag = strings.TrimSpace(part[:i])
			if qv := parseQ(part[i+1:]); qv >= 0 {
				q = qv
			}
		}

		// Match on the primary subtag: "es-ES" -> "es".
		primary := tag
		if i := strings.Index(primary, "-"); i != -1 {
			primary = primary[:i]
		}
		loc := Locale(strings.ToLower(primary))
		if !loc.Valid() {
			continue
		}
		if best == nil || q > best.q {
			best = &ranked{locale: loc, q: q}
		}
	}

	if best != nil {
		return best.locale
	}
	return Default
}

// parseQ extracts the numeric quality value from an Accept-Language parameter
// segment like "q=0.8". It returns -1 when no q parameter is present.
func parseQ(params string) float64 {
	for _, p := range strings.Split(params, ";") {
		p = strings.TrimSpace(p)
		if !strings.HasPrefix(p, "q=") {
			continue
		}
		if q, err := strconv.ParseFloat(strings.TrimPrefix(p, "q="), 64); err == nil {
			return q
		}
	}
	return -1
}
