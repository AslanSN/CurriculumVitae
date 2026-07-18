package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/internal/i18n"
	"github.com/a-h/templ"
)

type ExperienceStruct struct {
	Company, CompanyType, Contract, Position, RangeDate, ImageAlternative string
	Link, ImageSource                                                     templ.SafeURL
	Techs, Responsabilities, Extra                                        []string
}

// expNeutral holds the language-independent identity of a workplace. It is
// shared across every locale so links, dates, tech and images never drift.
type expNeutral struct {
	Company, RangeDate, ImageAlternative string
	ImageSource, Link                    templ.SafeURL
	Techs                                []string
}

// expProse holds the translatable copy for a workplace.
type expProse struct {
	CompanyType, Contract, Position string
	Responsabilities, Extra         []string
}

func buildExp(n expNeutral, p expProse) ExperienceStruct {
	return ExperienceStruct{
		Company:          n.Company,
		RangeDate:        n.RangeDate,
		ImageAlternative: n.ImageAlternative,
		ImageSource:      n.ImageSource,
		Link:             n.Link,
		Techs:            n.Techs,
		CompanyType:      p.CompanyType,
		Contract:         p.Contract,
		Position:         p.Position,
		Responsabilities: p.Responsabilities,
		Extra:            p.Extra,
	}
}

// --- neutral identity (shared across locales) -------------------------------

var (
	debosN = expNeutral{
		Company:          "Debos",
		RangeDate:        "10/24 - 08/26",
		ImageSource:      helpers.RepoURL + "/images/debos.svg",
		ImageAlternative: "Debos logo, the word debos in white on a dark tile",
		Link:             "https://debos.ai/",
		Techs:            []string{"React 19", "Next.js", "TypeScript", ".NET", "PostgreSQL", "Playwright", "Sentry", "Claude Code / MCP"},
	}
	nivimuN = expNeutral{
		Company:          "Nivimu",
		RangeDate:        "04/23 - 04/24",
		ImageSource:      helpers.RepoURL + "/images/nivimu.svg",
		ImageAlternative: "nivimu logo consists in a blue capital n with its name bellow",
		Link:             "https://nivimu.com/",
		Techs:            []string{"React", "Ant Design", "Emotion", "Redux", "Redux Saga", "TanStack (React) Query", "Mockoon", "Sentry", "bitDev", "GitLab"},
	}
	memorizameN = expNeutral{
		Company:          "Memorizame",
		RangeDate:        "12/22 - 04/23",
		ImageSource:      helpers.RepoURL + "/images/memorizame.webp",
		ImageAlternative: "memorizeme icon, three M in three different colors, one before another getting smaller",
		Link:             "",
		Techs:            []string{"Svelte", "Figma", "Miro", "Supabase", "TypeScript", "JavaScript", "Github"},
	}
	integroN = expNeutral{
		Company:          "Íntegro",
		RangeDate:        "02/22 - 04/23",
		ImageSource:      helpers.RepoURL + "/images/integro.webp",
		ImageAlternative: "integro writed in black with the o in blue",
		Link:             "",
		Techs:            []string{"Qwik", "Figma", "Illustrator", "TypeScript"},
	}
	attlosN = expNeutral{
		Company:          "Attlos",
		RangeDate:        "12/21 - 07/22",
		ImageSource:      helpers.RepoURL + "/images/attlos.webp",
		ImageAlternative: "Attlos logo",
		Link:             "https://www.youtube.com/channel/UC7hs7M2NfwWizyRIZkkOXVA",
		Techs:            []string{"Styled-Components", "TypeScript", "React", "Flux", "Redux", "Azure", "Git"},
	}
)

// --- translatable copy, per locale ------------------------------------------

var (
	debosEN = expProse{
		CompanyType: "Smart-building SaaS",
		Contract:    "Employee",
		Position:    "Full-Stack Engineer",
		Responsabilities: []string{
			"Owned frontend delivery; sustained it solo for ~8 months (3 large features to production, zero critical reverts)",
			"Took over the tech department's delivery coordination — deadlines, unblocking, comms — a function the CTO handed over; the team's delivery benchmark",
			"Built 6+ product modules end to end + a full-stack device search (React UI + .NET endpoint + PostgreSQL trigram indexes)",
			"Cross-cutting architectural refactors adopted as team conventions; built the design system and type-safe i18n across the app",
			"Built the end-to-end test suite (Playwright) from a dormant setup; added feature flags (Amplitude) + observability (Sentry)",
		},
		Extra: []string{
			"Authored the team's AI-engineering conventions (CLAUDE.md + versioned code-review skills)",
			"AI-augmented workflow: Claude Code, LLM agents, MCP",
		},
	}
	debosES = expProse{
		CompanyType: "SaaS de edificios inteligentes",
		Contract:    "Empleado",
		Position:    "Ingeniero Full-Stack",
		Responsabilities: []string{
			"Lideré la entrega del frontend; la sostuve en solitario durante ~8 meses (3 grandes funcionalidades a producción, cero reverts críticos)",
			"Asumí la coordinación de entrega del departamento técnico — plazos, desbloqueos, comunicación — una función que me delegó el CTO; el referente de entrega del equipo",
			"Construí más de 6 módulos de producto de punta a punta + un buscador de dispositivos full-stack (UI en React + endpoint en .NET + índices trigram en PostgreSQL)",
			"Refactors arquitectónicos transversales adoptados como convenciones del equipo; creé el design system y la i18n type-safe de toda la app",
			"Construí la suite de tests end-to-end (Playwright) desde una configuración inactiva; añadí feature flags (Amplitude) + observabilidad (Sentry)",
		},
		Extra: []string{
			"Redacté las convenciones de ingeniería con IA del equipo (CLAUDE.md + skills de code review versionadas)",
			"Flujo de trabajo potenciado por IA: Claude Code, agentes LLM, MCP",
		},
	}
	debosFR = expProse{
		CompanyType: "SaaS de bâtiments intelligents",
		Contract:    "Salarié",
		Position:    "Ingénieur Full-Stack",
		Responsabilities: []string{
			"Pilotage de la livraison du frontend ; maintenu seul pendant ~8 mois (3 grandes fonctionnalités en production, zéro revert critique)",
			"Pris en charge la coordination de la livraison du département technique — délais, déblocages, communication — une fonction déléguée par le CTO ; la référence de livraison de l'équipe",
			"Construit plus de 6 modules produit de bout en bout + une recherche d'appareils full-stack (UI React + endpoint .NET + index trigram PostgreSQL)",
			"Refactorisations architecturales transverses adoptées comme conventions d'équipe ; créé le design system et l'i18n type-safe de toute l'application",
			"Construit la suite de tests end-to-end (Playwright) à partir d'une configuration dormante ; ajouté des feature flags (Amplitude) + observabilité (Sentry)",
		},
		Extra: []string{
			"Rédigé les conventions d'ingénierie IA de l'équipe (CLAUDE.md + skills de code review versionnées)",
			"Flux de travail augmenté par l'IA : Claude Code, agents LLM, MCP",
		},
	}

	nivimuEN = expProse{
		CompanyType:      "Startup",
		Contract:         "Worker",
		Position:         "Frontend",
		Responsabilities: []string{"Code review", "Branch self-approval", "QA of UX", "Freedom of choice of tasks", "Deadlines"},
		Extra: []string{
			"Specializing in Time Zones, Ant Design and React Query",
			"Intense teamwork",
			"Cutting edge startup",
			"Complex application",
			"Short deadlines",
			"Agile, Lean, Extreme Programming",
		},
	}
	nivimuES = expProse{
		CompanyType:      "Startup",
		Contract:         "Trabajador",
		Position:         "Frontend",
		Responsabilities: []string{"Code review", "Autoaprobación de ramas", "QA de la UX", "Libertad para elegir tareas", "Plazos"},
		Extra: []string{
			"Especialización en zonas horarias, Ant Design y React Query",
			"Trabajo en equipo intenso",
			"Startup puntera",
			"Aplicación compleja",
			"Plazos ajustados",
			"Agile, Lean, Extreme Programming",
		},
	}
	nivimuFR = expProse{
		CompanyType:      "Startup",
		Contract:         "Salarié",
		Position:         "Frontend",
		Responsabilities: []string{"Code review", "Auto-approbation des branches", "QA de l'UX", "Liberté de choix des tâches", "Délais"},
		Extra: []string{
			"Spécialisation en fuseaux horaires, Ant Design et React Query",
			"Travail d'équipe intense",
			"Startup de pointe",
			"Application complexe",
			"Délais serrés",
			"Agile, Lean, Extreme Programming",
		},
	}

	memorizameEN = expProse{
		CompanyType:      "Project",
		Contract:         "Freelancer",
		Position:         "Full stack",
		Responsabilities: []string{"Design system", "UI/UX", "Architecture", "Layout", "Single frontend developer"},
		Extra:            []string{"App for learning with supermemo methodology"},
	}
	memorizameES = expProse{
		CompanyType:      "Proyecto",
		Contract:         "Freelance",
		Position:         "Full stack",
		Responsabilities: []string{"Design system", "UI/UX", "Arquitectura", "Maquetación", "Único desarrollador frontend"},
		Extra:            []string{"App para aprender con la metodología SuperMemo"},
	}
	memorizameFR = expProse{
		CompanyType:      "Projet",
		Contract:         "Freelance",
		Position:         "Full stack",
		Responsabilities: []string{"Design system", "UI/UX", "Architecture", "Intégration", "Seul développeur frontend"},
		Extra:            []string{"Application d'apprentissage avec la méthodologie SuperMemo"},
	}

	integroEN = expProse{
		CompanyType:      "Startup",
		Contract:         "Freelancer",
		Position:         "Full stack and more",
		Responsabilities: []string{"Design system", "UI/UX", "Architecture", "Layout", "Single software developer", "Branding", "Tech analyst"},
		Extra:            []string{"Thanks to my vision I have helped to create the identity of the Íntegro company"},
	}
	integroES = expProse{
		CompanyType:      "Startup",
		Contract:         "Freelance",
		Position:         "Full stack y más",
		Responsabilities: []string{"Design system", "UI/UX", "Arquitectura", "Maquetación", "Único desarrollador de software", "Branding", "Analista técnico"},
		Extra:            []string{"Con mi visión ayudé a crear la identidad de la empresa Íntegro"},
	}
	integroFR = expProse{
		CompanyType:      "Startup",
		Contract:         "Freelance",
		Position:         "Full stack et plus",
		Responsabilities: []string{"Design system", "UI/UX", "Architecture", "Intégration", "Seul développeur logiciel", "Branding", "Analyste technique"},
		Extra:            []string{"Grâce à ma vision, j'ai contribué à créer l'identité de l'entreprise Íntegro"},
	}

	attlosEN = expProse{
		CompanyType:      "Startup",
		Contract:         "Worker",
		Position:         "Front End",
		Responsabilities: []string{"Team Lead", "Deadlines", "Architecture"},
		Extra:            []string{"Implementing Tokenizations", "Agile Scrum"},
	}
	attlosES = expProse{
		CompanyType:      "Startup",
		Contract:         "Trabajador",
		Position:         "Front End",
		Responsabilities: []string{"Team Lead", "Plazos", "Arquitectura"},
		Extra:            []string{"Implementación de tokenizaciones", "Agile Scrum"},
	}
	attlosFR = expProse{
		CompanyType:      "Startup",
		Contract:         "Salarié",
		Position:         "Front End",
		Responsabilities: []string{"Team Lead", "Délais", "Architecture"},
		Extra:            []string{"Implémentation de tokenisations", "Agile Scrum"},
	}
)

// Workplaces is the experience list per locale, in reverse-chronological order.
var Workplaces = map[i18n.Locale][]ExperienceStruct{
	i18n.EN: {
		buildExp(debosN, debosEN),
		buildExp(nivimuN, nivimuEN),
		buildExp(memorizameN, memorizameEN),
		buildExp(integroN, integroEN),
		buildExp(attlosN, attlosEN),
	},
	i18n.ES: {
		buildExp(debosN, debosES),
		buildExp(nivimuN, nivimuES),
		buildExp(memorizameN, memorizameES),
		buildExp(integroN, integroES),
		buildExp(attlosN, attlosES),
	},
	i18n.FR: {
		buildExp(debosN, debosFR),
		buildExp(nivimuN, nivimuFR),
		buildExp(memorizameN, memorizameFR),
		buildExp(integroN, integroFR),
		buildExp(attlosN, attlosFR),
	},
}
