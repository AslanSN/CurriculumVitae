package i18n

import "strings"

// SitemapXML returns a sitemap that lists every locale's canonical URL with
// reciprocal hreflang alternates, following Google's guidance for
// internationalized sites (each URL declares all language variants).
func SitemapXML() string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">` + "\n")
	for _, loc := range Supported {
		b.WriteString("  <url>\n")
		b.WriteString("    <loc>" + loc.CanonicalURL() + "</loc>\n")
		for _, alt := range Supported {
			b.WriteString(`    <xhtml:link rel="alternate" hreflang="` + alt.Lang() + `" href="` + alt.CanonicalURL() + `"/>` + "\n")
		}
		b.WriteString(`    <xhtml:link rel="alternate" hreflang="x-default" href="` + EN.CanonicalURL() + `"/>` + "\n")
		b.WriteString("  </url>\n")
	}
	b.WriteString("</urlset>\n")
	return b.String()
}

// RobotsTXT allows all crawlers and points at the sitemap.
func RobotsTXT() string {
	return "User-agent: *\nAllow: /\n\nSitemap: " + SiteURL + "/sitemap.xml\n"
}
