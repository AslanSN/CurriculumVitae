package constants

import "github.com/a-h/templ"

type Challenge struct {
	Company, Title, State, Dates, Description string
	Duration                                  int
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
		RepoLink:    "https://github.com/AslanSN/Logitravel-Technical-Proof",
		AppLink:     "https://logitravel-technical-proof.vercel.app/",
		Description: "",
	}
	ArupChallenge = Challenge{
		Company:     "Arup",
		Title:       "Arup Technical Test",
		State:       "Not enough experience",
		Dates:       "From 7 to 14 June 2022",
		Duration:    7,
		RepoLink:    "https://github.com/AslanSN/arup-technicalTest",
		AppLink:     "https://arup-technical-test-deploy.vercel.app/",
		Description: "",
	}
	NivimuChallenge = Challenge{
		Company:     "Nivimu",
		Title:       "Nivimu Code Test",
		State:       "HIRED",
		Dates:       "From 20 to 21 March 2023",
		Duration:    1,
		RepoLink:    "https://github.com/AslanSN/nivimu-front-codetest",
		AppLink:     "https://aslansn.github.io/nivimu-front-codetest/",
		Description: "",
	}
	InditexChallenge = Challenge{
		Company:     "Inditex",
		Title:       "Zara Marvel Web Challenge",
		State:       "Process stopped prior to review of my technical proof",
		Dates:       "From 1 to 8 August 2024",
		Duration:    7,
		RepoLink:    "https://github.com/AslanSN/zara-web-challenge",
		AppLink:     "https://zara-marvelous.vercel.app/",
		Description: "",
	}
)
