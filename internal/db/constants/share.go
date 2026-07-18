package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/internal/i18n"
	"github.com/a-h/templ"
)

// Message is the pre-filled text for the social share buttons, per locale.
var Message = map[i18n.Locale]string{
	i18n.EN: `🌟 Check out Alan Staub Negro's portfolio 🌟
👉 Take a look here: aslansn-cv.vercel.app
If you like what you see, consider sharing it! It's great to support talented developers and creatives. 💖
`,
	i18n.ES: `🌟 Descubre el portfolio de Alan Staub Negro 🌟
👉 Échale un vistazo aquí: aslansn-cv.vercel.app
Si te gusta lo que ves, ¡compártelo! Apoyar a desarrolladores y creativos con talento siempre suma. 💖
`,
	i18n.FR: `🌟 Découvrez le portfolio d'Alan Staub Negro 🌟
👉 Jetez-y un œil ici : aslansn-cv.vercel.app
Si vous aimez ce que vous voyez, partagez-le ! Soutenir les développeurs et créatifs de talent, c'est toujours précieux. 💖
`,
}

type SocialShareButton struct {
	Icon helpers.IconLabelParams
	URL  string
}

var SocialShareButtons = []SocialShareButton{
	Facebook,
	Linkedin,
	XTwiter,
	Whatsapp,
}

var (
	Facebook = SocialShareButton{
		Icon: helpers.IconLabelParams{
			Id:          "facebook",
			Label:       "",
			Source:      templ.SafeURL(helpers.RepoURL + "/icons/facebook.svg"),
			Alternative: "Facebook Icon",
		},
		URL: "https://www.facebook.com/sharer/sharer.php?quote={message}",
	}
	Linkedin = SocialShareButton{
		Icon: helpers.IconLabelParams{
				Id:          "linkedin",
				Label:       "",
				Source:      templ.SafeURL(helpers.RepoURL + "/icons/linkedin.svg"),
				Alternative: "Linkedin Icon",
			
		},
		URL:  "https://www.linkedin.com/shareArticle?mini=true&title={message}",
	}
	XTwiter = SocialShareButton{
		Icon: helpers.IconLabelParams{
			Id:          "xtwitter",
			Label:       "",
			Source:      templ.SafeURL(helpers.RepoURL + "/icons/x.svg"),
			Alternative: "X or Twiter Icon",
		},
		URL:  "https://x.com/intent/tweet?text={message}",
	}
	Whatsapp = SocialShareButton{
		Icon: helpers.IconLabelParams{
			Id:          "whatsapp",
			Label:       "",
			Source:      templ.SafeURL(helpers.RepoURL + "/icons/whatsapp.svg"),
			Alternative: "Whatsapp Icon",
		},
		URL:  "https://api.whatsapp.com/send?text={message}",
	}
)
