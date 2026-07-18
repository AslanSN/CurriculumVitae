package constants

import (
	"github.com/AslanSN/CurriculumVitae/internal/i18n"
	"github.com/a-h/templ"
)

type Challenge struct {
	Company, Title, State, Dates, Description string
	Duration                                  int
	Highlight                                 bool
	RepoLink                                  templ.SafeURL
	AppLink                                   templ.SafeURL
}

// challengeNeutral holds the language-independent facts of a technical test.
type challengeNeutral struct {
	Company   string
	Duration  int
	Highlight bool
	RepoLink  templ.SafeURL
	AppLink   templ.SafeURL
}

// challengeProse holds the translatable copy for a technical test.
type challengeProse struct {
	Title, State, Dates, Description string
}

func buildChallenge(n challengeNeutral, p challengeProse) Challenge {
	return Challenge{
		Company:     n.Company,
		Duration:    n.Duration,
		Highlight:   n.Highlight,
		RepoLink:    n.RepoLink,
		AppLink:     n.AppLink,
		Title:       p.Title,
		State:       p.State,
		Dates:       p.Dates,
		Description: p.Description,
	}
}

// --- neutral facts (shared across locales) ----------------------------------

var (
	inditexN = challengeNeutral{
		Company:  "Inditex",
		Duration: 7,
		RepoLink: "https://github.com/AslanSN/zara-web-challenge",
		AppLink:  "",
	}
	nivimuChN = challengeNeutral{
		Company:   "Nivimu",
		Duration:  1,
		Highlight: true,
		RepoLink:  "https://github.com/AslanSN/nivimu-front-codetest",
		AppLink:   "https://aslansn.github.io/nivimu-front-codetest/",
	}
	arupN = challengeNeutral{
		Company:  "Arup",
		Duration: 7,
		RepoLink: "https://github.com/AslanSN/arup-technicalTest",
		AppLink:  "https://arup-technical-test-deploy.vercel.app/",
	}
	logitravelN = challengeNeutral{
		Company:   "Logitravel",
		Duration:  3,
		Highlight: true,
		RepoLink:  "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:   "https://logitravel-technical-proof.vercel.app/",
	}
)

// --- translatable copy, per locale ------------------------------------------

var (
	inditexEN = challengeProse{
		Title:       "Zara Marvel Web Challenge",
		State:       "Process stopped prior to review of my technical proof",
		Dates:       "From 1 to 8 August 2024",
		Description: "Marvel character explorer (list, search, detail, favorites) in Next.js + TypeScript. State via Context + useReducer (no Redux), favorites persisted to localStorage, Vitest tests and Husky pre-commit hooks.",
	}
	inditexES = challengeProse{
		Title:       "Zara Marvel Web Challenge",
		State:       "Proceso detenido antes de revisar mi prueba técnica",
		Dates:       "Del 1 al 8 de agosto de 2024",
		Description: "Explorador de personajes de Marvel (lista, búsqueda, detalle, favoritos) en Next.js + TypeScript. Estado con Context + useReducer (sin Redux), favoritos persistidos en localStorage, tests con Vitest y hooks pre-commit con Husky.",
	}
	inditexFR = challengeProse{
		Title:       "Zara Marvel Web Challenge",
		State:       "Processus interrompu avant l'examen de mon test technique",
		Dates:       "Du 1er au 8 août 2024",
		Description: "Explorateur de personnages Marvel (liste, recherche, détail, favoris) en Next.js + TypeScript. État via Context + useReducer (sans Redux), favoris persistés dans localStorage, tests Vitest et hooks pre-commit avec Husky.",
	}

	nivimuChEN = challengeProse{
		Title:       "Nivimu Code Test",
		State:       "HIRED",
		Dates:       "From 20 to 21 March 2023",
		Description: "Filterable, sortable user-data table with a live summary card that re-syncs to the top row as sorting and filtering change. React + TypeScript + Redux Toolkit + Ant Design — delivered in a single day.",
	}
	nivimuChES = challengeProse{
		Title:       "Prueba de código de Nivimu",
		State:       "CONTRATADO",
		Dates:       "Del 20 al 21 de marzo de 2023",
		Description: "Tabla de datos de usuarios filtrable y ordenable con una tarjeta de resumen en tiempo real que se resincroniza con la primera fila al cambiar el orden y los filtros. React + TypeScript + Redux Toolkit + Ant Design — entregada en un solo día.",
	}
	nivimuChFR = challengeProse{
		Title:       "Test de code Nivimu",
		State:       "EMBAUCHÉ",
		Dates:       "Du 20 au 21 mars 2023",
		Description: "Tableau de données utilisateurs filtrable et triable avec une carte de résumé en temps réel qui se resynchronise avec la première ligne au fil des tris et filtres. React + TypeScript + Redux Toolkit + Ant Design — livré en une seule journée.",
	}

	arupEN = challengeProse{
		Title:       "Arup Technical Test",
		State:       "Not enough experience",
		Dates:       "From 7 to 14 June 2022",
		Description: "Construction-site communications manager: a three-panel UI — discipline/status filter sidebar, sortable questions list and detail view. React (Flux) + SASS, SOLID and documented components.",
	}
	arupES = challengeProse{
		Title:       "Prueba técnica de Arup",
		State:       "Experiencia insuficiente",
		Dates:       "Del 7 al 14 de junio de 2022",
		Description: "Gestor de comunicaciones de obra: una UI de tres paneles — barra lateral de filtros por disciplina/estado, lista ordenable de preguntas y vista de detalle. React (Flux) + SASS, SOLID y componentes documentados.",
	}
	arupFR = challengeProse{
		Title:       "Test technique Arup",
		State:       "Expérience insuffisante",
		Dates:       "Du 7 au 14 juin 2022",
		Description: "Gestionnaire de communications de chantier : une UI à trois panneaux — barre latérale de filtres par discipline/statut, liste de questions triable et vue de détail. React (Flux) + SASS, SOLID et composants documentés.",
	}

	logitravelEN = challengeProse{
		Title:       "Logitravel - PrimeIT Technical Proof",
		State:       "Candidate with best technical skills",
		Dates:       "From 1 to 4 July 2022",
		Description: "Text-list manager with add / select / delete and full undo–redo. React + Redux Toolkit + styled-components, a centralized design-token system and SOLID structure — built in a 3-day proof.",
	}
	logitravelES = challengeProse{
		Title:       "Logitravel - Prueba técnica de PrimeIT",
		State:       "Candidato con las mejores competencias técnicas",
		Dates:       "Del 1 al 4 de julio de 2022",
		Description: "Gestor de listas de texto con añadir / seleccionar / borrar y deshacer–rehacer completo. React + Redux Toolkit + styled-components, un sistema centralizado de design tokens y estructura SOLID — construido en una prueba de 3 días.",
	}
	logitravelFR = challengeProse{
		Title:       "Logitravel - Test technique PrimeIT",
		State:       "Candidat aux meilleures compétences techniques",
		Dates:       "Du 1er au 4 juillet 2022",
		Description: "Gestionnaire de listes de texte avec ajouter / sélectionner / supprimer et annuler–rétablir complet. React + Redux Toolkit + styled-components, un système centralisé de design tokens et une structure SOLID — construit lors d'un test de 3 jours.",
	}
)

// Challenges is the technical-test list per locale.
var Challenges = map[i18n.Locale][]Challenge{
	i18n.EN: {
		buildChallenge(inditexN, inditexEN),
		buildChallenge(nivimuChN, nivimuChEN),
		buildChallenge(arupN, arupEN),
		buildChallenge(logitravelN, logitravelEN),
	},
	i18n.ES: {
		buildChallenge(inditexN, inditexES),
		buildChallenge(nivimuChN, nivimuChES),
		buildChallenge(arupN, arupES),
		buildChallenge(logitravelN, logitravelES),
	},
	i18n.FR: {
		buildChallenge(inditexN, inditexFR),
		buildChallenge(nivimuChN, nivimuChFR),
		buildChallenge(arupN, arupFR),
		buildChallenge(logitravelN, logitravelFR),
	},
}
