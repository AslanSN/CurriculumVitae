package i18n

import "context"

// Dict is the type-safe catalog of UI strings. Every user-facing string that
// lives in a template is a field here — a mistyped field name is a compile
// error, and dict_test.go guarantees every locale fills every field.
type Dict struct {
	// Global chrome
	SkipToContent string
	NavHome       string
	Language      string // language-switcher label

	// Section labels — used both in the nav and as <h2> headings
	SectionAbout      string
	SectionSkills     string
	SectionExperience string
	SectionChallenges string
	SectionArtifact   string

	// Actions
	Share      string
	Contact    string
	ContactMe  string // Contact button + contact-modal title
	DownloadCV string

	// Hero
	HeroRole    string
	HeroTagline string

	// Challenges section
	ChallengesIntro     string
	ChallengeRepository string
	ChallengeLiveDemo   string
	Day                 string
	Days                string

	// Artifact section (the self-directed, open-source flagship)
	ArtifactIntro    string
	ArtifactColophon string

	// Experience-card labels
	Stack      string
	WhatIDid   string
	Highlights string

	// Colophon (the "Go + templ + Tailwind CSS," prefix stays literal)
	ColophonHeading string
	ColophonBody    string
	SourceOnGitHub  string

	// Share modal
	ShareTitle string

	// Contact — pre-filled outreach (email subject/body, LinkedIn & WhatsApp)
	ContactSubject string
	ContactMessage string

	// Accessibility labels
	Close     string
	OpenMenu  string
	CopyPhone string
	CopyEmail string

	// <head> metadata
	MetaTitle       string
	MetaDescription string
}

// dicts holds one fully-populated Dict per supported locale.
var dicts = map[Locale]Dict{
	EN: {
		SkipToContent: "Skip to content",
		NavHome:       "Home",
		Language:      "Language",

		SectionAbout:      "About",
		SectionSkills:     "Skills",
		SectionExperience: "Experience",
		SectionChallenges: "The verdict",
		SectionArtifact:   "The artifact",

		Share:      "Share",
		Contact:    "Contact",
		ContactMe:  "Contact me",
		DownloadCV: "Download CV",

		HeroRole:    "Senior Full-Stack Engineer",
		HeroTagline: "AI-native · React / Next.js + .NET · platform",

		ChallengesIntro:     "Company technical tests under a deadline — how I approach an unfamiliar problem, and the verdict each one earned.",
		ChallengeRepository: "Repository",
		ChallengeLiveDemo:   "Live demo",
		Day:                 "day",
		Days:                "days",

		ArtifactIntro:    "Self-directed and open source. Not the word “AI-native” — the thing itself: a harness that makes an AI produce correct code, one you can clone and run.",
		ArtifactColophon: "How it was built",

		Stack:      "Stack",
		WhatIDid:   "What I did",
		Highlights: "Highlights",

		ColophonHeading: "How this site is built",
		ColophonBody:    "server-rendered on Vercel — no client framework, just fast HTML. The stack is a deliberate choice, not my daily React/.NET: a content site doesn't need a SPA, so I fit the tool to the problem — type-safe end to end, down to an i18n catalog that fails the build if a translation is missing. The honest cost: a smaller ecosystem than React. The AI-native workflow behind it — Claude Code — is how I actually ship, day to day.",
		SourceOnGitHub:  "Source on GitHub",

		ShareTitle: "Share this portfolio",

		ContactSubject: "First contact",
		ContactMessage: "Hi Alan, I came across your portfolio and would like to get in touch.",

		Close:     "Close",
		OpenMenu:  "Open menu",
		CopyPhone: "Copy phone number",
		CopyEmail: "Copy email",

		MetaTitle:       "Alan Staub Negro — Senior Full-Stack Engineer (AI-native)",
		MetaDescription: "Senior full-stack engineer (~4.5y): React 19 / Next.js + TypeScript on the front, real .NET / PostgreSQL on the back, AI-native (Claude Code, MCP). Trilingual, remote-first.",
	},

	ES: {
		SkipToContent: "Saltar al contenido",
		NavHome:       "Inicio",
		Language:      "Idioma",

		SectionAbout:      "Sobre mí",
		SectionSkills:     "Competencias",
		SectionExperience: "Experiencia",
		SectionChallenges: "El veredicto",
		SectionArtifact:   "El artefacto",

		Share:      "Compartir",
		Contact:    "Contacto",
		ContactMe:  "Contáctame",
		DownloadCV: "Descargar CV",

		HeroRole:    "Ingeniero Full-Stack Senior",
		HeroTagline: "AI-native · React / Next.js + .NET · plataforma",

		ChallengesIntro:     "Pruebas técnicas de empresa contra reloj — cómo abordo un problema desconocido, y el veredicto que se llevó cada una.",
		ChallengeRepository: "Repositorio",
		ChallengeLiveDemo:   "Demo en vivo",
		Day:                 "día",
		Days:                "días",

		ArtifactIntro:    "Autodirigido y de código abierto. No la palabra «AI-native» — la cosa en sí: un harness que hace que una IA produzca código correcto, y que puedes clonar y ejecutar.",
		ArtifactColophon: "Cómo se construyó",

		Stack:      "Stack",
		WhatIDid:   "Qué hice",
		Highlights: "Destacados",

		ColophonHeading: "Cómo está hecha esta web",
		ColophonBody:    "renderizada en el servidor y desplegada en Vercel — sin framework de cliente, solo HTML rápido. El stack es una elección deliberada, no mi React/.NET del día a día: un sitio de contenido no necesita una SPA, así que ajusté la herramienta al problema — tipado de punta a punta, hasta un catálogo i18n que rompe el build si falta una traducción. El coste honesto: un ecosistema menor que el de React. El flujo AI-native que hay detrás — Claude Code — es como trabajo de verdad, cada día.",
		SourceOnGitHub:  "Código en GitHub",

		ShareTitle: "Comparte este portfolio",

		ContactSubject: "Primer contacto",
		ContactMessage: "Hola Alan, he visto tu portfolio y me gustaría ponerme en contacto contigo.",

		Close:     "Cerrar",
		OpenMenu:  "Abrir menú",
		CopyPhone: "Copiar número de teléfono",
		CopyEmail: "Copiar correo",

		MetaTitle:       "Alan Staub Negro — Ingeniero Full-Stack Senior (AI-native)",
		MetaDescription: "Ingeniero full-stack senior (~4,5 años): React 19 / Next.js + TypeScript en el front, .NET / PostgreSQL real en el back, AI-native (Claude Code, MCP). Trilingüe, remote-first.",
	},

	FR: {
		SkipToContent: "Aller au contenu",
		NavHome:       "Accueil",
		Language:      "Langue",

		SectionAbout:      "À propos",
		SectionSkills:     "Compétences",
		SectionExperience: "Expérience",
		SectionChallenges: "Le verdict",
		SectionArtifact:   "L'artefact",

		Share:      "Partager",
		Contact:    "Contact",
		ContactMe:  "Me contacter",
		DownloadCV: "Télécharger le CV",

		HeroRole:    "Ingénieur Full-Stack Senior",
		HeroTagline: "AI-native · React / Next.js + .NET · plateforme",

		ChallengesIntro:     "Tests techniques d'entreprise sous contrainte de délai — comment j'aborde un problème inconnu, et le verdict obtenu pour chacun.",
		ChallengeRepository: "Dépôt",
		ChallengeLiveDemo:   "Démo en ligne",
		Day:                 "jour",
		Days:                "jours",

		ArtifactIntro:    "En autonomie et open source. Pas le mot « AI-native » — la chose elle-même : un harness qui fait qu'une IA produit du code correct, que vous pouvez cloner et exécuter.",
		ArtifactColophon: "Comment il est construit",

		Stack:      "Stack",
		WhatIDid:   "Ce que j'ai fait",
		Highlights: "Points forts",

		ColophonHeading: "Comment ce site est construit",
		ColophonBody:    "rendu côté serveur et déployé sur Vercel — sans framework côté client, juste du HTML rapide. Le stack est un choix délibéré, pas mon React/.NET du quotidien : un site de contenu n'a pas besoin d'une SPA, alors j'ai adapté l'outil au problème — typé de bout en bout, jusqu'à un catalogue i18n qui casse le build s'il manque une traduction. Le coût honnête : un écosystème plus petit que React. Le flux de travail AI-native derrière — Claude Code — est ma façon de livrer, au quotidien.",
		SourceOnGitHub:  "Code source sur GitHub",

		ShareTitle: "Partagez ce portfolio",

		ContactSubject: "Premier contact",
		ContactMessage: "Bonjour Alan, j'ai vu votre portfolio et j'aimerais entrer en contact avec vous.",

		Close:     "Fermer",
		OpenMenu:  "Ouvrir le menu",
		CopyPhone: "Copier le numéro de téléphone",
		CopyEmail: "Copier l'e-mail",

		MetaTitle:       "Alan Staub Negro — Ingénieur Full-Stack Senior (AI-native)",
		MetaDescription: "Ingénieur full-stack senior (~4,5 ans) : React 19 / Next.js + TypeScript côté front, .NET / PostgreSQL réel côté back, AI-native (Claude Code, MCP). Trilingue, remote-first.",
	},
}

// T returns the string catalog for the locale carried by ctx.
func T(ctx context.Context) Dict { return dicts[FromContext(ctx)] }

// Of returns the string catalog for an explicit locale.
func Of(l Locale) Dict { return dicts[l] }
