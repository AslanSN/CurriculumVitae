package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/a-h/templ"
)

type SocialMediaData struct {
	helpers.IconLabelParams
	SocialMediaLink templ.SafeURL
}

var SocialMediaList = []SocialMediaData{
	LinkedIn,
	Github,
	X,
}

var AslanImageAlt string = "Foto en primer plano de Alan Staub Negro sonriendo"

var (
	LinkedIn = SocialMediaData{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "linkedIn",
			Label:       "",
			Source:      helpers.RepoURL + "/icons/linkedin.svg",
			Alternative: "LinkedIn Icon",
		},
		SocialMediaLink: "https://www.linkedin.com/in/alanstaubnegro/",
	}
	Github = SocialMediaData{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "github",
			Label:       "",
			Source:      helpers.RepoURL + "/icons/github.svg",
			Alternative: "Github Icon",
		},
		SocialMediaLink: "https://github.com/AslanSN",
	}
	X = SocialMediaData{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "x",
			Label:       "",
			Source:      helpers.RepoURL + "/icons/x.svg",
			Alternative: "X, before twitter, icon",
		},
		SocialMediaLink: "https://x.com/AslanSN_",
	}
)
