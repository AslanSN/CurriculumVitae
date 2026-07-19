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
