package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/i18n"
	"github.com/a-h/templ"
)

type ExperienceStruct struct {
	Company, CompanyType, Contract, Position, RangeDate, ImageAlternative string
	Link, ImageSource                                                     templ.SafeURL
	Techs, Responsabilities, Extra                                        []string
	// Compact renders the workplace as a single achievement line (older,
	// less-central roles) instead of the full Stack / What-I-did card.
	Compact bool
}

// expNeutral holds the language-independent identity of a workplace. It is
// shared across every locale so links, dates, tech and images never drift.
type expNeutral struct {
	Company, RangeDate, ImageAlternative string
	ImageSource, Link                    templ.SafeURL
	Techs                                []string
	Compact                              bool
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
		Compact:          n.Compact,
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
		Techs:            []string{"React", "TypeScript", "Ant Design", "Emotion", "Redux", "Redux Saga", "TanStack Query", "Bit", "Mockoon", "Sentry"},
	}
	memorizameN = expNeutral{
		Company:          "Memorizame",
		RangeDate:        "12/22 - 04/23",
		ImageSource:      helpers.RepoURL + "/images/memorizame.webp",
		ImageAlternative: "memorizeme icon, three M in three different colors, one before another getting smaller",
		Link:             "",
		Techs:            []string{"Svelte", "Supabase", "TypeScript", "Figma"},
		Compact:          true,
	}
	integroN = expNeutral{
		Company:          "Íntegro",
		RangeDate:        "02/22 - 04/23",
		ImageSource:      helpers.RepoURL + "/images/integro.webp",
		ImageAlternative: "integro writed in black with the o in blue",
		Link:             "",
		Techs:            []string{"Qwik", "TypeScript", "Figma"},
		Compact:          true,
	}
	attlosN = expNeutral{
		Company:          "Attlos",
		RangeDate:        "12/21 - 07/22",
		ImageSource:      helpers.RepoURL + "/images/attlos.webp",
		ImageAlternative: "Attlos logo",
		Link:             "https://www.youtube.com/channel/UC7hs7M2NfwWizyRIZkkOXVA",
		Techs:            []string{"TypeScript", "React", "Redux", "Azure"},
		Compact:          true,
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
		CompanyType: "HR PWA startup",
		Contract:    "Employee",
		Position:    "Frontend Developer",
		Responsabilities: []string{
			"Owned features end to end with high individual responsibility — code review, branch self-approval, UX QA",
			"Led the JavaScript → TypeScript migration across a large-scale HR PWA",
			"Built an internal component library, design system and micro-frontends (Bit)",
			"Handled time-zone-correct scheduling across the product",
		},
	}
	nivimuES = expProse{
		CompanyType: "Startup de PWA de RR. HH.",
		Contract:    "Empleado",
		Position:    "Desarrollador Frontend",
		Responsabilities: []string{
			"Ownership de features de principio a fin con alta responsabilidad individual — code review, autoaprobación de ramas, QA de UX",
			"Lideré la migración JavaScript → TypeScript en una PWA de RR. HH. a gran escala",
			"Construí una biblioteca de componentes interna, design system y micro-frontends (Bit)",
			"Gestioné el manejo correcto de zonas horarias en todo el producto",
		},
	}
	nivimuFR = expProse{
		CompanyType: "Startup PWA RH",
		Contract:    "Salarié",
		Position:    "Développeur Frontend",
		Responsabilities: []string{
			"Ownership des fonctionnalités de bout en bout avec forte responsabilité individuelle — code review, auto-approbation des branches, QA de l'UX",
			"Piloté la migration JavaScript → TypeScript sur une PWA RH à grande échelle",
			"Construit une bibliothèque de composants interne, un design system et des micro-frontends (Bit)",
			"Assuré la gestion correcte des fuseaux horaires dans tout le produit",
		},
	}

	memorizameEN = expProse{
		CompanyType:      "Freelance project",
		Contract:         "Freelancer",
		Position:         "Full stack",
		Responsabilities: []string{"Sole frontend for a spaced-repetition learning app (SuperMemo method) — design system, UI/UX and architecture."},
	}
	memorizameES = expProse{
		CompanyType:      "Proyecto freelance",
		Contract:         "Freelance",
		Position:         "Full stack",
		Responsabilities: []string{"Único frontend de una app de aprendizaje por repetición espaciada (método SuperMemo) — design system, UI/UX y arquitectura."},
	}
	memorizameFR = expProse{
		CompanyType:      "Projet freelance",
		Contract:         "Freelance",
		Position:         "Full stack",
		Responsabilities: []string{"Seul frontend d'une application d'apprentissage par répétition espacée (méthode SuperMemo) — design system, UI/UX et architecture."},
	}

	integroEN = expProse{
		CompanyType:      "Startup",
		Contract:         "Freelancer",
		Position:         "Full stack",
		Responsabilities: []string{"Sole developer and technical analyst for a startup — company site end to end, design system and branding."},
	}
	integroES = expProse{
		CompanyType:      "Startup",
		Contract:         "Freelance",
		Position:         "Full stack",
		Responsabilities: []string{"Único desarrollador y analista técnico de una startup — web de empresa de principio a fin, design system y branding."},
	}
	integroFR = expProse{
		CompanyType:      "Startup",
		Contract:         "Freelance",
		Position:         "Full stack",
		Responsabilities: []string{"Seul développeur et analyste technique d'une startup — site d'entreprise de bout en bout, design system et branding."},
	}

	attlosEN = expProse{
		CompanyType:      "Startup",
		Contract:         "Intern",
		Position:         "Front End",
		Responsabilities: []string{"First professional TypeScript/React role — built product features in an Agile/Scrum team."},
	}
	attlosES = expProse{
		CompanyType:      "Startup",
		Contract:         "Becario",
		Position:         "Front End",
		Responsabilities: []string{"Primer puesto profesional con TypeScript/React — desarrollé features de producto en un equipo Agile/Scrum."},
	}
	attlosFR = expProse{
		CompanyType:      "Startup",
		Contract:         "Stagiaire",
		Position:         "Front End",
		Responsabilities: []string{"Premier poste professionnel en TypeScript/React — développé des fonctionnalités produit dans une équipe Agile/Scrum."},
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
