package constants

import (
	iconLabel "github.com/AslanSN/CV/components/iconLabel"
	"github.com/AslanSN/CV/db"
	"github.com/a-h/templ"
)

type SocialMediaData struct {
	iconLabel.IconLabelParams
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
		IconLabelParams: iconLabel.IconLabelParams{
			Id:          "linkedIn",
			Label:       "LinkedIn",
			Source:      db.RepoURL + "/icons/linkedin.svg",
			Alternative: "LinkedIn Icon",
		},
		SocialMediaLink: "https://www.linkedin.com/in/alanstaubnegro/",
	}
	Github = SocialMediaData{
		IconLabelParams: iconLabel.IconLabelParams{
			Id:          "github",
			Label:       "GitHub",
			Source:      db.RepoURL + "/icons/github.svg",
			Alternative: "Github Icon",
		},
		SocialMediaLink: "https://github.com/AslanSN",
	}
	X = SocialMediaData{
		IconLabelParams: iconLabel.IconLabelParams{
			Id:          "x",
			Label:       "X",
			Source:      db.RepoURL + "/icons/x-social-media-round-icon.svg",
			Alternative: "X, before twitter, icon",
		},
		SocialMediaLink: "https://x.com/AslanSN_",
	}
)
