package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/a-h/templ"
)

var Message = `🌟 Check out the impressive portfolio of AslanSN 🌟
👉 Take a look here: aslansn-cv.vercel.app 
If you like what you see, consider sharing it! It's great to support talented developers and creatives. 💖
`

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
