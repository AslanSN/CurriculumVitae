package i18n

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResolve(t *testing.T) {
	cases := []struct {
		name         string
		target       string
		acceptLang   string
		cookie       string
		wantLocale   Locale
		wantRedirect string
	}{
		{"es path", "/es", "", "", ES, ""},
		{"fr path", "/fr", "", "", FR, ""},
		{"es subpath", "/es/whatever", "", "", ES, ""},
		{"root no signals", "/", "", "", EN, ""},
		{"root accept-language es redirects", "/", "es-ES,es;q=0.9", "", ES, "/es"},
		{"root accept-language fr redirects", "/", "fr-FR,fr;q=0.9", "", FR, "/fr"},
		{"query overrides accept-language", "/?lang=en", "es-ES,es;q=0.9", "", EN, ""},
		{"query fr redirects to clean path", "/?lang=fr", "", "", FR, "/fr"},
		{"cookie fr redirects", "/", "", "fr", FR, "/fr"},
		{"cookie en pins english", "/", "es-ES,es;q=0.9", "en", EN, ""},
		{"path wins over cookie", "/es", "", "en", ES, ""},
		{"unsupported language falls back to en", "/", "de-DE,de;q=0.9", "", EN, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, c.target, nil)
			if c.acceptLang != "" {
				r.Header.Set("Accept-Language", c.acceptLang)
			}
			if c.cookie != "" {
				r.AddCookie(&http.Cookie{Name: cookieName, Value: c.cookie})
			}
			gotLocale, gotRedirect := Resolve(r)
			if gotLocale != c.wantLocale || gotRedirect != c.wantRedirect {
				t.Errorf("Resolve(target=%q, AL=%q, cookie=%q) = (%q, %q), want (%q, %q)",
					c.target, c.acceptLang, c.cookie, gotLocale, gotRedirect, c.wantLocale, c.wantRedirect)
			}
		})
	}
}

func TestSitemapAndRobots(t *testing.T) {
	sm := SitemapXML()
	for _, loc := range Supported {
		if !strings.Contains(sm, "<loc>"+loc.CanonicalURL()+"</loc>") {
			t.Errorf("sitemap missing <loc> for %q", loc)
		}
		if !strings.Contains(sm, `hreflang="`+loc.Lang()+`"`) {
			t.Errorf("sitemap missing hreflang for %q", loc)
		}
	}
	if !strings.Contains(sm, `hreflang="x-default"`) {
		t.Error("sitemap missing x-default alternate")
	}
	if robots := RobotsTXT(); !strings.Contains(robots, SiteURL+"/sitemap.xml") {
		t.Error("robots.txt missing sitemap URL")
	}
}
