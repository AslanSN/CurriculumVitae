package views

// revealScript is a tiny, dependency-free scroll-reveal enhancer emitted early in
// <head> (rendered raw via templ.Raw). Progressive enhancement: it hides content
// (via the .reveal-ready class) only when IntersectionObserver is available AND the
// visitor has not requested reduced motion — so no-JS and reduced-motion visitors
// always see everything. Each .reveal element is revealed once, when it enters view.
const revealScript = `(function () {
  var root = document.documentElement;
  if (!('IntersectionObserver' in window) ||
      !window.matchMedia('(prefers-reduced-motion: no-preference)').matches) return;
  root.classList.add('reveal-ready');
  window.addEventListener('DOMContentLoaded', function () {
    var io = new IntersectionObserver(function (entries) {
      entries.forEach(function (entry) {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible');
          io.unobserve(entry.target);
        }
      });
    }, { rootMargin: '0px 0px -8% 0px', threshold: 0.08 });
    document.querySelectorAll('.reveal').forEach(function (node) { io.observe(node); });
  });
})();`

// revealScriptTag wraps revealScript in a <script> element. templ v0.2.x does NOT
// evaluate @-expressions inside a literal <script> element (it emits them as text),
// so we render the whole tag via templ.Raw(revealScriptTag) as an expression child.
const revealScriptTag = "<script>" + revealScript + "</script>"

// personLDJSON is the schema.org Person structured data emitted in <head>
// (rendered raw via templ.Raw). Keep values in sync with the CV.
const personLDJSON = `{
  "@context": "https://schema.org",
  "@type": "Person",
  "name": "Alan Staub Negro",
  "url": "https://aslansn-cv.vercel.app/",
  "image": "https://aslansn-cv.vercel.app/assets/images/AslanSNPhoto.webp",
  "jobTitle": "Senior Full-Stack Engineer",
  "email": "mailto:aslan.staub@pm.me",
  "address": {
    "@type": "PostalAddress",
    "addressLocality": "Palma de Mallorca",
    "addressCountry": "ES"
  },
  "knowsLanguage": ["es", "fr", "en"],
  "knowsAbout": ["React", "Next.js", "TypeScript", ".NET", "PostgreSQL", "Node.js", "Go", "AI engineering", "Claude Code", "Web performance"],
  "sameAs": [
    "https://github.com/AslanSN",
    "https://www.linkedin.com/in/alanstaubnegro/",
    "https://x.com/AslanSN_"
  ]
}`

// personLDJSONTag wraps personLDJSON in its <script> element, rendered via
// templ.Raw (see revealScriptTag for why a literal <script> element won't render
// the @-expression). Fixes the structured data that was previously emitted as the
// literal text "@templ.Raw(personLDJSON)".
const personLDJSONTag = `<script type="application/ld+json">` + personLDJSON + `</script>`

// themeScript resolves and applies the colour theme BEFORE first paint (it is
// emitted synchronously at the top of <head>), so there's no flash of the wrong
// theme. Order of precedence: the visitor's explicit choice (localStorage) →
// otherwise the OS/browser preference. It writes data-theme="light|dark" on
// <html>; the light palette in output.css keys off :root[data-theme="light"].
// It also exposes window.__toggleTheme() for the toggle button, and keeps the
// page in sync with the OS while the visitor hasn't chosen explicitly.
const themeScript = `(function () {
  var root = document.documentElement, KEY = 'theme';
  function stored() { try { return localStorage.getItem(KEY); } catch (e) { return null; } }
  function resolve() {
    var s = stored();
    if (s === 'light' || s === 'dark') return s;
    return (window.matchMedia && window.matchMedia('(prefers-color-scheme: light)').matches) ? 'light' : 'dark';
  }
  root.setAttribute('data-theme', resolve());
  window.__toggleTheme = function () {
    var next = root.getAttribute('data-theme') === 'light' ? 'dark' : 'light';
    root.setAttribute('data-theme', next);
    try { localStorage.setItem(KEY, next); } catch (e) {}
  };
  if (window.matchMedia) {
    try {
      window.matchMedia('(prefers-color-scheme: light)').addEventListener('change', function (e) {
        if (!stored()) root.setAttribute('data-theme', e.matches ? 'light' : 'dark');
      });
    } catch (e) {}
  }
})();`

// themeScriptTag wraps themeScript in a <script> element (rendered via templ.Raw,
// see revealScriptTag). Emitted as early as possible in <head> to avoid a flash.
const themeScriptTag = "<script>" + themeScript + "</script>"

// noscriptThemeTag is the no-JS fallback: with scripting off, data-theme is never
// set, so honour the OS preference directly. Covers the core palette tokens (the
// same values output.css uses); component-level refinements are JS-only.
const noscriptThemeTag = `<noscript><style>@media (prefers-color-scheme: light){:root{color-scheme:light;` +
	`--color-base-950:#f6f2ea;--color-base-900:#fefcf8;--color-base-800:#efe9de;` +
	`--color-accent:#0e7490;--color-white:#241f19;` +
	`--color-zinc-100:#1c1a17;--color-zinc-200:#2b2823;--color-zinc-300:#413d36;` +
	`--color-zinc-400:#6b655b;--color-zinc-500:#736d63;}}</style></noscript>`
