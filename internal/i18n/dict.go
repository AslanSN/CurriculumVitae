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
		SectionChallenges: "Challenges",

		Share:      "Share",
		Contact:    "Contact",
		ContactMe:  "Contact me",
		DownloadCV: "Download CV",

		HeroRole:    "Senior Full-Stack Engineer",
		HeroTagline: "AI-native · React / Next.js + .NET · platform",

		ChallengesIntro:     "Company technical tests — how I approach an unfamiliar problem under a deadline, and how each one landed.",
		ChallengeRepository: "Repository",
		ChallengeLiveDemo:   "Live demo",
		Day:                 "day",
		Days:                "days",

		Stack:      "Stack",
		WhatIDid:   "What I did",
		Highlights: "Highlights",

		ColophonHeading: "How this site is built",
		ColophonBody:    "server-rendered and deployed on Vercel — no client-side framework, just fast HTML. Rebuilt with an AI-native workflow (Claude Code), the way I ship day to day.",
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
		SectionChallenges: "Pruebas técnicas",

		Share:      "Compartir",
		Contact:    "Contacto",
		ContactMe:  "Contáctame",
		DownloadCV: "Descargar CV",

		HeroRole:    "Ingeniero Full-Stack Senior",
		HeroTagline: "AI-native · React / Next.js + .NET · plataforma",

		ChallengesIntro:     "Pruebas técnicas de empresas — cómo abordo un problema desconocido contra reloj, y en qué acabó cada una.",
		ChallengeRepository: "Repositorio",
		ChallengeLiveDemo:   "Demo en vivo",
		Day:                 "día",
		Days:                "días",

		Stack:      "Stack",
		WhatIDid:   "Qué hice",
		Highlights: "Destacados",

		ColophonHeading: "Cómo está hecha esta web",
		ColophonBody:    "renderizada en el servidor y desplegada en Vercel — sin framework de cliente, solo HTML rápido. Reconstruida con un flujo AI-native (Claude Code), tal como trabajo cada día.",
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
		SectionChallenges: "Tests techniques",

		Share:      "Partager",
		Contact:    "Contact",
		ContactMe:  "Me contacter",
		DownloadCV: "Télécharger le CV",

		HeroRole:    "Ingénieur Full-Stack Senior",
		HeroTagline: "AI-native · React / Next.js + .NET · plateforme",

		ChallengesIntro:     "Tests techniques d'entreprises — comment j'aborde un problème inconnu sous contrainte de délai, et le résultat de chacun.",
		ChallengeRepository: "Dépôt",
		ChallengeLiveDemo:   "Démo en ligne",
		Day:                 "jour",
		Days:                "jours",

		Stack:      "Stack",
		WhatIDid:   "Ce que j'ai fait",
		Highlights: "Points forts",

		ColophonHeading: "Comment ce site est construit",
		ColophonBody:    "rendu côté serveur et déployé sur Vercel — sans framework côté client, juste du HTML rapide. Reconstruit avec un flux de travail AI-native (Claude Code), comme je livre au quotidien.",
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
