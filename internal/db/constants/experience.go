package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/a-h/templ"
)

type ExperienceStruct struct {
	Company, CompanyType, Contract, Position, RangeDate, ImageAlternative string
	Link, ImageSource                                                     templ.SafeURL
	Techs, Responsabilities, Extra                                        []string
}

var Workplaces = []ExperienceStruct{
	Debos,
	Nivimu,
	Memorizame,
	Integro,
	Worksut,
}

var (
	Debos = ExperienceStruct{
		Company:          "Debos",
		CompanyType:      "Smart-building SaaS",
		Contract:         "Employee",
		Position:         "Full-Stack Engineer",
		RangeDate:        "10/24 - 08/26",
		ImageSource:      helpers.RepoURL + "/images/debos.svg",
		ImageAlternative: "Debos logo, the word debos in white on a dark tile",
		Link:             "https://debos.ai/",
		Techs: []string{
			"React 19",
			"Next.js",
			"TypeScript",
			".NET",
			"PostgreSQL",
			"Playwright",
			"Sentry",
			"Claude Code / MCP",
		},
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
	Nivimu = ExperienceStruct{
		Company:          "Nivimu",
		CompanyType:      "Startup",
		Contract:         "Worker",
		Position:         "Frontend",
		RangeDate:        "04/23 - 04/24",
		ImageSource:      helpers.RepoURL + "/images/nivimu.svg",
		ImageAlternative: "nivimu logo consists in a blue capital n with its name bellow",
		Link:             "https://nivimu.com/",
		Techs: []string{
			"React",
			"Ant Design",
			"Emotion",
			"Redux",
			"Redux Saga",
			"TanStack (React) Query",
			"Mockoon",
			"Sentry",
			"bitDev",
			"GitLab",
		},
		Responsabilities: []string{
			"Code review",
			"Branch self-approval",
			"QA of UX",
			"Freedom of choice of tasks",
			"Deadlines",
		},
		Extra: []string{
			"Specializing in Time Zones, Ant Design and React Query",
			"Intense teamwork",
			"Cutting edge startup",
			"Complex application",
			"Short deadlines",
			"Agile, Lean, Extreme Programming",
		},
	}
	Memorizame = ExperienceStruct{
		Company:          "Memorizame",
		CompanyType:      "Project",
		Contract:         "Freelancer",
		Position:         "Full stack",
		RangeDate:        "12/22 - 04/23",
		ImageSource:      helpers.RepoURL + "/images/memorizame.webp",
		ImageAlternative: "memorizeme icon, three M in three different colors, one before another getting smaller",
		Link:             "",
		Techs: []string{
			"Svelte",
			"Figma",
			"Miro",
			"Supabase",
			"TypeScript",
			"JavaScript",
			"Github",
		},
		Responsabilities: []string{
			"Design system",
			"UI/UX",
			"Architecture",
			"Layout",
			"Single frontend developer",
		},
		Extra: []string{
			"App for learning with supermemo methodology",
		},
	}
	Integro = ExperienceStruct{
		Company:          "Íntegro",
		CompanyType:      "Startup",
		Contract:         "Freelancer",
		Position:         "Full stack and more",
		RangeDate:        "02/22 - 04/23",
		ImageSource:      helpers.RepoURL + "/images/integro.webp",
		ImageAlternative: "integro writed in black with the o in blue",
		Link:             "",
		Techs: []string{
			"Qwik",
			"Figma",
			"Illustrator",
			"TypeScript",
		},
		Responsabilities: []string{
			"Design system",
			"UI/UX",
			"Architecture",
			"Layout",
			"Single software developer",
			"Branding",
			"Tech analyst",
		},
		Extra: []string{
			"Thanks to my vision I have helped to create the identity of the Íntegro company",
		},
	}
	Worksut = ExperienceStruct{
		Company:          "Attlos",
		CompanyType:      "Startup",
		Contract:         "Worker",
		Position:         "Front End",
		RangeDate:        "12/21 - 07/22",
		ImageSource:      helpers.RepoURL + "/images/attlos.webp",
		ImageAlternative: "Attlos logo",

		Link: "https://www.youtube.com/channel/UC7hs7M2NfwWizyRIZkkOXVA",
		Techs: []string{
			"Styled-Components",
			"TypeScript",
			"React",
			"Flux",
			"Redux",
			"Azure",
			"Git",
		},
		Responsabilities: []string{
			"Team Lead",
			"Deadlines",
			"Architecture",
		},
		Extra: []string{
			"Implementing Tokenizations",
			"Agile Scrum",
		},
	}
)
