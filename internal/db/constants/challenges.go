package constants

import "github.com/a-h/templ"

type Challenge struct {
	Company, Title, State, Dates, Description string
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
		Description: "",
		State:       "Candidate with best technical skills",
		Dates:       "From 1 to 4 July 2022",
		RepoLink:    "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:     templ.SafeURL(""),
	}
	ArupChallenge = Challenge{
		Company:     "Arup",
		Title:       "Arup Technical Test",
		State:       "Not enough experience",
		Description: "",
		Dates:       "From 7 to 14 June 2022",
		RepoLink:    "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:     templ.SafeURL(""),
	}
	NivimuChallenge = Challenge{
		Company:     "Nivimu",
		Title:       "Nivimu Code Test",
		State:       "HIRED",
		Description: "",
		Dates:       "From 20 to 21 March 2023",
		RepoLink:    "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:     "https://aslansn.github.io/nivimu-front-codetest/",
	}
	InditexChallenge = Challenge{
		Company:     "Inditex",
		Title:       "Zara Marvel Web Challenge",
		State:       "Process stopped prior to review of my technical proof",
		Description: "",
		Dates:       "From 1 to 8 August 2024",
		RepoLink:    "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:     "https://zara-marvelous.vercel.app/",
	}
)