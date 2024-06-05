package constants

import (
	iconLabels "github.com/AslanSN/CV/components/iconLabel"
	"github.com/AslanSN/CV/db"
	"github.com/a-h/templ"
)

type Skill struct {
	iconLabels.IconLabelParams
	Link templ.SafeURL
}

var SkillsList = []Skill{
	JavaScript,
	TypeScript,
	React,
	Redux,
	TanStack,
	AntDesign,
	Go,
}

var (
	JavaScript = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "javaScript",
			Label:       "JavaScript",
			Alternative: "JavaScript Icon, white J S letters, yellow background",
			Source:      db.RepoURL + "/icons/javascript.svg",
		},
		Link: "https://es.wikipedia.org/wiki/JavaScript",
	}
	TypeScript = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "typeScript",
			Label:       "TypeScript",
			Alternative: "TypeScript Icon, white T S letters, blue background",
			Source:      db.RepoURL + "/icons/typescript.svg",
		},
		Link: "https://www.typescriptlang.org/",
	}
	React = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "react",
			Label:       "React",
			Alternative: "React Icon, an atom illustration in blue color",
			Source:      db.RepoURL + "/icons/react.svg",
		},
		Link: "https://react.dev/",
	}
	Redux = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "redux",
			Label:       "Redux",
			Alternative: "Redux icon, a purple triquetra with dots",
			Source:      db.RepoURL + "/icons/redux.svg",
		},
		Link: "https://redux.js.org/",
	}
	TanStack = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "reactQuery",
			Label:       `TanStack Query`,
			Alternative: "TanStack query old icon that is as a flower with yellow center and red petals",
			Source:      db.RepoURL + "/icons/reactQuery.svg",
		},
		Link: "https://tanstack.com/query/latest/docs/framework/react/overview",
	}
	AntDesign = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "antD",
			Label:       "Ant Design",
			Alternative: "Ant design logo, a sided square with three quartes blue lines, the rest in red and a red dot at the center",
			Source:      "https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg",
		},
		Link: "https://ant.design/",
	}
	Go = Skill{
		IconLabelParams: iconLabels.IconLabelParams{
			Id:          "go",
			Label:       "Go",
			Alternative: "Golang logo is the GO word in blue color",
			Source:      db.RepoURL + "/icons/golang.svg",
		},
		Link: "https://go.dev/",
	}
)
