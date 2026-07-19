package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/i18n"
	"github.com/a-h/templ"
)

// Artifact is a self-directed, open-source project shown on its own — the
// executable proof behind the "AI-native" label rather than the label itself.
type Artifact struct {
	Name         string // brand name, not localized
	RepoLink     templ.SafeURL
	ColophonLink templ.SafeURL
	VideoMP4     templ.SafeURL
	VideoPoster  templ.SafeURL
	Eyebrow      string
	Tagline      string
	Lead         string
	VideoAlt     string
	VideoCaption string
	Highlights   []string
}

// artifactNeutral holds the language-independent facts: links and media.
type artifactNeutral struct {
	Name         string
	RepoLink     templ.SafeURL
	ColophonLink templ.SafeURL
	VideoMP4     templ.SafeURL
	VideoPoster  templ.SafeURL
}

// artifactProse holds the translatable copy for the featured project.
type artifactProse struct {
	Eyebrow      string
	Tagline      string
	Lead         string
	VideoAlt     string
	VideoCaption string
	Highlights   []string
}

func buildArtifact(n artifactNeutral, p artifactProse) Artifact {
	return Artifact{
		Name:         n.Name,
		RepoLink:     n.RepoLink,
		ColophonLink: n.ColophonLink,
		VideoMP4:     n.VideoMP4,
		VideoPoster:  n.VideoPoster,
		Eyebrow:      p.Eyebrow,
		Tagline:      p.Tagline,
		Lead:         p.Lead,
		VideoAlt:     p.VideoAlt,
		VideoCaption: p.VideoCaption,
		Highlights:   p.Highlights,
	}
}

// --- neutral facts (shared across locales) ----------------------------------

var gotchaN = artifactNeutral{
	Name:         "Gotcha",
	RepoLink:     "https://github.com/AslanSN/gotcha",
	ColophonLink: "https://github.com/AslanSN/gotcha/blob/main/COLOPHON.md",
	// Recorded from gotcha's own robust-search slice; re-encoded H.264 (~115 KB,
	// down from a 1.6 MB GIF) so the demo doesn't tax the "just fast HTML" budget.
	VideoMP4:    helpers.RepoURL + "/images/gotcha-robust-search.mp4",
	VideoPoster: helpers.RepoURL + "/images/gotcha-robust-search.jpg",
}

// --- translatable copy, per locale ------------------------------------------

var (
	gotchaEN = artifactProse{
		Eyebrow: "Open source · MCP · evals",
		Tagline: "An AI-native engineering harness — and the catalog of subtle bugs that compile, pass review, and ship anyway.",
		Lead:    "The gap that matters isn't using an LLM to type faster — it's engineering the system that makes it produce correct code in a domain full of traps. Gotcha is that system: a defect catalog, versioned review skills, an MCP server that hands those rules to an agent while it writes, and an eval suite that measures whether the AI trips each trap.",
		VideoAlt:     "Screen recording: typing filters a ticket search over ~500k rows, a superseded request is cancelled, and matching results stream in.",
		VideoCaption: "Flagship — robust ticket search over ~500k rows: min-length, debounce and request cancellation on the client; a trigram scan that stays a BitmapOr, not a Seq Scan, on the server.",
		Highlights: []string{
			"16-defect catalog",
			"MCP server · 5 agent tools",
			"Evals 11/16 · recall & precision 11/11",
			".NET · PostgreSQL · Next.js / React 19",
		},
	}
	gotchaES = artifactProse{
		Eyebrow: "Código abierto · MCP · evals",
		Tagline: "Un harness de ingeniería AI-native — y el catálogo de bugs sutiles que compilan, pasan la revisión y llegan a producción igualmente.",
		Lead:    "La diferencia que importa no es usar un LLM para teclear más rápido — es diseñar el sistema que le hace producir código correcto en un dominio lleno de trampas. Gotcha es ese sistema: un catálogo de defectos, skills de revisión versionadas, un servidor MCP que le pasa esas reglas a un agente mientras escribe, y una suite de evals que mide si la IA cae en cada trampa.",
		VideoAlt:     "Grabación de pantalla: al teclear se filtra un buscador de tickets sobre ~500k filas, se cancela una petición obsoleta y aparecen los resultados coincidentes.",
		VideoCaption: "Flagship — búsqueda robusta de tickets sobre ~500k filas: mínimo de caracteres, debounce y cancelación de peticiones en el cliente; una búsqueda trigram que sigue siendo un BitmapOr, no un Seq Scan, en el servidor.",
		Highlights: []string{
			"Catálogo de 16 defectos",
			"Servidor MCP · 5 herramientas",
			"Evals 11/16 · recall y precisión 11/11",
			".NET · PostgreSQL · Next.js / React 19",
		},
	}
	gotchaFR = artifactProse{
		Eyebrow: "Open source · MCP · evals",
		Tagline: "Un harness d'ingénierie AI-native — et le catalogue de bugs subtils qui compilent, passent la revue et partent en production quand même.",
		Lead:    "L'écart qui compte n'est pas d'utiliser un LLM pour taper plus vite — c'est de concevoir le système qui lui fait produire du code correct dans un domaine plein de pièges. Gotcha est ce système : un catalogue de défauts, des skills de revue versionnées, un serveur MCP qui fournit ces règles à un agent pendant qu'il écrit, et une suite d'evals qui mesure si l'IA tombe dans chaque piège.",
		VideoAlt:     "Capture d'écran : la saisie filtre une recherche de tickets sur ~500k lignes, une requête obsolète est annulée et les résultats correspondants s'affichent.",
		VideoCaption: "Flagship — recherche robuste de tickets sur ~500k lignes : minimum de caractères, debounce et annulation des requêtes côté client ; une recherche trigram qui reste un BitmapOr, pas un Seq Scan, côté serveur.",
		Highlights: []string{
			"Catalogue de 16 défauts",
			"Serveur MCP · 5 outils",
			"Evals 11/16 · rappel et précision 11/11",
			".NET · PostgreSQL · Next.js / React 19",
		},
	}
)

// Artifacts is the featured self-directed project, per locale.
var Artifacts = map[i18n.Locale]Artifact{
	i18n.EN: buildArtifact(gotchaN, gotchaEN),
	i18n.ES: buildArtifact(gotchaN, gotchaES),
	i18n.FR: buildArtifact(gotchaN, gotchaFR),
}
