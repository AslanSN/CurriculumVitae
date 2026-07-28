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
	ID           string // ASCII slug for the section anchor / list id
	RepoLink     templ.SafeURL
	ColophonLink templ.SafeURL
	LiveLink     templ.SafeURL // live-demo URL; when set it becomes the primary CTA
	// EmbedURL renders a live, interactive <iframe> preview in place of a video.
	// Exactly one of EmbedURL / VideoMP4 should be set per artifact.
	EmbedURL     templ.SafeURL
	VideoMP4     templ.SafeURL
	VideoPoster  templ.SafeURL
	Eyebrow      string
	Tagline      string
	Lead         string
	VideoAlt     string // also the <iframe> title when EmbedURL is used
	VideoCaption string
	Highlights   []string
}

// artifactNeutral holds the language-independent facts: links and media.
type artifactNeutral struct {
	Name         string
	ID           string
	RepoLink     templ.SafeURL
	ColophonLink templ.SafeURL
	LiveLink     templ.SafeURL
	EmbedURL     templ.SafeURL
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
		ID:           n.ID,
		RepoLink:     n.RepoLink,
		ColophonLink: n.ColophonLink,
		LiveLink:     n.LiveLink,
		EmbedURL:     n.EmbedURL,
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
	ID:           "gotcha",
	RepoLink:     "https://github.com/AslanSN/gotcha",
	ColophonLink: "https://github.com/AslanSN/gotcha/blob/main/COLOPHON.md",
	// Recorded from gotcha's own robust-search slice; re-encoded H.264 (~115 KB,
	// down from a 1.6 MB GIF) so the demo doesn't tax the "just fast HTML" budget.
	VideoMP4:    helpers.RepoURL + "/images/gotcha-robust-search.mp4",
	VideoPoster: helpers.RepoURL + "/images/gotcha-robust-search.jpg",
}

// Memorízame is the design-engineering artifact: a spaced-repetition study PWA
// Alan designed and built end to end. The product he did this design work in is
// private, so this is a public, self-contained demo (github.com/AslanSN/memorizame-v2)
// that carries the design evidence. Shown as a LIVE embed rather than a video.
var memorizameArtN = artifactNeutral{
	Name:         "Memorízame",
	ID:           "memorizame",
	RepoLink:     "https://github.com/AslanSN/memorizame-v2",
	LiveLink:     "https://memorizame-v2.vercel.app/",
	EmbedURL:     "https://memorizame-v2.vercel.app/",
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

	memorizameArtEN = artifactProse{
		Eyebrow: "Design engineering · open source",
		Tagline: "A spaced-repetition study app I designed and built end to end — brand, design system and interface, grounded in the science of learning.",
		Lead:    "The product I did this design work in is private, so I rebuilt it as an open, installable demo you can actually use — seeded with public-domain content. Every surface is mine: the two-tone logo, a tokenised light/dark design system (WCAG-AA / APCA-checked colour, three typefaces chosen for reading science), a trilingual UI and a real SM-2 scheduler. Held to the same senior bar as the rest — a design-token contract test and an axe-core accessibility gate fail the build on any regression.",
		VideoAlt:     "Live embedded demo of Memorízame, a spaced-repetition study PWA",
		VideoCaption: "Live and interactive — three modes (study, review, test), a real SM-2 scheduler, trilingual UI in light and dark. Try it right here.",
		Highlights: []string{
			"Design system · tokens",
			"SM-2 · offline PWA",
			"WCAG 2.2 AA · APCA colour",
			"Trilingual · light/dark",
			"SvelteKit 2 · Svelte 5",
		},
	}
	memorizameArtES = artifactProse{
		Eyebrow: "Ingeniería de diseño · código abierto",
		Tagline: "Una app de estudio por repetición espaciada que diseñé y construí de punta a punta — marca, design system e interfaz, fundamentados en la ciencia del aprendizaje.",
		Lead:    "El producto donde hice este trabajo de diseño es privado, así que lo reconstruí como una demo abierta e instalable que puedes usar de verdad — con contenido de dominio público. Cada superficie es mía: el logo de dos tonos, un design system tokenizado claro/oscuro (color verificado WCAG-AA / APCA, tres tipografías elegidas por la ciencia de la lectura), una interfaz trilingüe y un planificador SM-2 real. Al mismo nivel senior que el resto — un test de contrato de design tokens y un gate de accesibilidad (axe-core) rompen el build ante cualquier regresión.",
		VideoAlt:     "Demo embebido en vivo de Memorízame, una PWA de estudio por repetición espaciada",
		VideoCaption: "En vivo e interactivo — tres modos (estudiar, repasar, test), un planificador SM-2 real, interfaz trilingüe en claro y oscuro. Pruébalo aquí mismo.",
		Highlights: []string{
			"Design system · tokens",
			"SM-2 · PWA offline",
			"Color WCAG 2.2 AA · APCA",
			"Trilingüe · claro/oscuro",
			"SvelteKit 2 · Svelte 5",
		},
	}
	memorizameArtFR = artifactProse{
		Eyebrow: "Ingénierie du design · open source",
		Tagline: "Une app d'étude en répétition espacée que j'ai conçue et construite de bout en bout — identité, design system et interface, fondés sur la science de l'apprentissage.",
		Lead:    "Le produit où j'ai réalisé ce travail de design est privé ; je l'ai donc reconstruit en une démo ouverte et installable, réellement utilisable — avec du contenu du domaine public. Chaque surface est de moi : le logo bicolore, un design system tokenisé clair/sombre (couleurs vérifiées WCAG-AA / APCA, trois typographies choisies selon la science de la lecture), une interface trilingue et un vrai planificateur SM-2. Tenu au même niveau senior que le reste — un test de contrat de design tokens et un gate d'accessibilité (axe-core) cassent le build à la moindre régression.",
		VideoAlt:     "Démo intégrée en direct de Memorízame, une PWA d'étude en répétition espacée",
		VideoCaption: "En direct et interactif — trois modes (étudier, réviser, test), un vrai planificateur SM-2, interface trilingue en clair et sombre. Essayez-le ici.",
		Highlights: []string{
			"Design system · tokens",
			"SM-2 · PWA hors-ligne",
			"Couleur WCAG 2.2 AA · APCA",
			"Trilingue · clair/sombre",
			"SvelteKit 2 · Svelte 5",
		},
	}
)

// Artifacts lists the self-directed projects per locale, in feature order:
// Gotcha (AI-native engineering) first, then Memorízame (design engineering).
var Artifacts = map[i18n.Locale][]Artifact{
	i18n.EN: {
		buildArtifact(gotchaN, gotchaEN),
		buildArtifact(memorizameArtN, memorizameArtEN),
	},
	i18n.ES: {
		buildArtifact(gotchaN, gotchaES),
		buildArtifact(memorizameArtN, memorizameArtES),
	},
	i18n.FR: {
		buildArtifact(gotchaN, gotchaFR),
		buildArtifact(memorizameArtN, memorizameArtFR),
	},
}
