package constants

import "github.com/a-h/templ"

type Challenge struct {
	Company, Title, State, Dates, Description string
	Duration                                  int
	Highlight                                 bool
	RepoLink                                  templ.SafeURL
	AppLink                                   templ.SafeURL
}

var Challenges = []Challenge{
	InditexChallenge,
	NivimuChallenge,
	ArupChallenge,
	LogitravelChallenge,
}

var (
	LogitravelChallenge = Challenge{
		Company:     "Logitravel",
		Title:       "Logitravel - PrimeIT Technical Proof",
		State:       "Candidate with best technical skills",
		Dates:       "From 1 to 4 July 2022",
		Duration:    3,
		Highlight:   true,
		RepoLink:    "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:     "https://logitravel-technical-proof.vercel.app/",
		Description: "Text-list manager with add / select / delete and full undo–redo. React + Redux Toolkit + styled-components, a centralized design-token system and SOLID structure — built in a 3-day proof.",
	}
	ArupChallenge = Challenge{
		Company:     "Arup",
		Title:       "Arup Technical Test",
		State:       "Not enough experience",
		Dates:       "From 7 to 14 June 2022",
		Duration:    7,
		RepoLink:    "https://github.com/AslanSN/arup-technicalTest",
		AppLink:     "https://arup-technical-test-deploy.vercel.app/",
		Description: "Construction-site communications manager: a three-panel UI — discipline/status filter sidebar, sortable questions list and detail view. React (Flux) + SASS, SOLID and documented components.",
	}
	NivimuChallenge = Challenge{
		Company:     "Nivimu",
		Title:       "Nivimu Code Test",
		State:       "HIRED",
		Dates:       "From 20 to 21 March 2023",
		Duration:    1,
		Highlight:   true,
		RepoLink:    "https://github.com/AslanSN/nivimu-front-codetest",
		AppLink:     "https://aslansn.github.io/nivimu-front-codetest/",
		Description: "Filterable, sortable user-data table with a live summary card that re-syncs to the top row as sorting and filtering change. React + TypeScript + Redux Toolkit + Ant Design — delivered in a single day.",
	}
	InditexChallenge = Challenge{
		Company:     "Inditex",
		Title:       "Zara Marvel Web Challenge",
		State:       "Process stopped prior to review of my technical proof",
		Dates:       "From 1 to 8 August 2024",
		Duration:    7,
		RepoLink:    "https://github.com/AslanSN/zara-web-challenge",
		AppLink:     "",
		Description: "Marvel character explorer (list, search, detail, favorites) in Next.js + TypeScript. State via Context + useReducer (no Redux), favorites persisted to localStorage, Vitest tests and Husky pre-commit hooks.",
	}
)
