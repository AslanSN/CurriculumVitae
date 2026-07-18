package constants

import (
	"github.com/AslanSN/CurriculumVitae/i18n"
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
// Ordered by outcome (the strongest third-party verdict first), not by date.

var (
	debosChN = challengeNeutral{
		Company:   "Debos",
		Duration:  0, // live interview, not a timed take-home
		Highlight: true,
	}
	nivimuChN = challengeNeutral{
		Company:   "Nivimu",
		Duration:  1,
		Highlight: true,
		RepoLink:  "https://github.com/AslanSN/nivimu-front-codetest",
		AppLink:   "https://aslansn.github.io/nivimu-front-codetest/",
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
	debosChEN = challengeProse{
		Title:       "Debos — Live problem-solving interview",
		State:       "Hired on the spot",
		Dates:       "October 2024",
		Description: "Two-part technical interview; the real test was a design problem — a search box hitting a backend with highly variable latency: how do you handle it? My answer, three layers: a minimum query length, a 200–500 ms debounce, and an AbortController cancelling the in-flight request so a slower earlier response can't overwrite a newer one — the classic typeahead race condition. Hired on the spot; joined as an external consultant, then moved in-house. I later shipped exactly this in production — a full-stack device search (React + .NET + PostgreSQL trigram indexes).",
	}
	debosChES = challengeProse{
		Title:       "Debos — Entrevista de resolución de problemas",
		State:       "Contratado en el acto",
		Dates:       "Octubre de 2024",
		Description: "Entrevista técnica en dos partes; la de verdad era un problema de diseño — un buscador que ataca un backend de latencia muy variable: ¿cómo lo gestionas? Mi respuesta, tres capas: un mínimo de caracteres, un debounce de 200–500 ms y un AbortController que cancela la petición en vuelo para que una respuesta anterior más lenta no sobrescriba a una más nueva — la clásica race condition de un typeahead. Contratado en el acto; entré como consultor externo y luego pasé a plantilla. Después construí exactamente esto en producción — un buscador de dispositivos full-stack (React + .NET + índices trigram en PostgreSQL).",
	}
	debosChFR = challengeProse{
		Title:       "Debos — Entretien de résolution de problèmes",
		State:       "Embauché sur-le-champ",
		Dates:       "Octobre 2024",
		Description: "Entretien technique en deux parties ; le vrai test était un problème de conception — un champ de recherche interrogeant un backend à latence très variable : comment le gérer ? Ma réponse, trois couches : un minimum de caractères, un debounce de 200–500 ms et un AbortController qui annule la requête en cours pour qu'une réponse antérieure plus lente n'écrase pas une plus récente — la classique race condition d'un typeahead. Embauché sur-le-champ ; arrivé comme consultant externe, puis passé en interne. J'ai ensuite construit exactement cela en production — une recherche d'appareils full-stack (React + .NET + index trigram PostgreSQL).",
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

	logitravelEN = challengeProse{
		Title:       "Logitravel - PrimeIT Technical Proof",
		State:       "Best technical candidate",
		Dates:       "From 1 to 4 July 2022",
		Description: "Text-list manager with add / select / delete and full undo–redo. React + Redux Toolkit + styled-components, a centralized design-token system and SOLID structure — built in a 3-day proof.",
	}
	logitravelES = challengeProse{
		Title:       "Logitravel - Prueba técnica de PrimeIT",
		State:       "Mejor candidato técnico",
		Dates:       "Del 1 al 4 de julio de 2022",
		Description: "Gestor de listas de texto con añadir / seleccionar / borrar y deshacer–rehacer completo. React + Redux Toolkit + styled-components, un sistema centralizado de design tokens y estructura SOLID — construido en una prueba de 3 días.",
	}
	logitravelFR = challengeProse{
		Title:       "Logitravel - Test technique PrimeIT",
		State:       "Meilleur candidat technique",
		Dates:       "Du 1er au 4 juillet 2022",
		Description: "Gestionnaire de listes de texte avec ajouter / sélectionner / supprimer et annuler–rétablir complet. React + Redux Toolkit + styled-components, un système centralisé de design tokens et une structure SOLID — construit lors d'un test de 3 jours.",
	}
)

// Challenges is the technical-test list per locale, ordered by outcome:
// Debos (hired on the spot) → Nivimu (hired) → Logitravel (best technical candidate).
var Challenges = map[i18n.Locale][]Challenge{
	i18n.EN: {
		buildChallenge(debosChN, debosChEN),
		buildChallenge(nivimuChN, nivimuChEN),
		buildChallenge(logitravelN, logitravelEN),
	},
	i18n.ES: {
		buildChallenge(debosChN, debosChES),
		buildChallenge(nivimuChN, nivimuChES),
		buildChallenge(logitravelN, logitravelES),
	},
	i18n.FR: {
		buildChallenge(debosChN, debosChFR),
		buildChallenge(nivimuChN, nivimuChFR),
		buildChallenge(logitravelN, logitravelFR),
	},
}
