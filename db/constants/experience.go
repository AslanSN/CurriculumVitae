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
	Nivimu,
	Memorizame,
	Integro,
	Worksut,
}

var (
	Nivimu = ExperienceStruct{
		Company:          "Nivimu",
		CompanyType:      "Startup",
		Contract:         "Worker",
		Position:         "Frontend",
		RangeDate:        "04/23 - 04/24",
		ImageSource:      "https://nivimu.com/wp-content/uploads/2021/05/logo-nivimu.svg",
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
		ImageAlternative: "nivimu logo consists in a blue capital n with its name bellow",

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
