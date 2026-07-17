package views

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
