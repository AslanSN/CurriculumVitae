package constants

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/i18n"
	"github.com/a-h/templ"
)

type Skill struct {
	helpers.IconLabelParams
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
	Node,
	MySQL,
}

type SkillGroup struct {
	Name  string
	Items []string
}

// frontendItems and alsoItems are pure tech names — identical across locales,
// so they're shared to avoid drift.
var (
	frontendItems = []string{"React 19", "Next.js (App Router)", "TypeScript", "TanStack Query", "React Hook Form", "Zod", "Zustand", "Tailwind CSS", "Radix UI"}
	alsoItems     = []string{"Go", "Svelte / SvelteKit", "React Native / Expo", "Python / Flask", "Figma (UX/UI)"}
)

// SkillGroups is the grouped skill list per locale. Group names, the AI-related
// phrasings and the spoken languages are translated; tool/library names are not.
var SkillGroups = map[i18n.Locale][]SkillGroup{
	i18n.EN: {
		{Name: "AI & agentic dev", Items: []string{"Claude Code", "LLM agents (Claude, GPT/Codex)", "MCP", "AI-assisted code review", "Team AI conventions (CLAUDE.md)"}},
		{Name: "Frontend", Items: frontendItems},
		{Name: "Backend & platform", Items: []string{".NET (C#)", "EF Core", "Node.js", "PostgreSQL", "REST APIs", "CI/CD (GitHub Actions)", "Vercel", "Sentry", "Amplitude", "Playwright", "Vitest"}},
		{Name: "Also", Items: alsoItems},
		{Name: "Languages", Items: []string{"Spanish (native)", "French (bilingual)", "English (professional)"}},
	},
	i18n.ES: {
		{Name: "IA y desarrollo agéntico", Items: []string{"Claude Code", "Agentes LLM (Claude, GPT/Codex)", "MCP", "Code review asistido por IA", "Convenciones de IA de equipo (CLAUDE.md)"}},
		{Name: "Frontend", Items: frontendItems},
		{Name: "Backend y plataforma", Items: []string{".NET (C#)", "EF Core", "Node.js", "PostgreSQL", "APIs REST", "CI/CD (GitHub Actions)", "Vercel", "Sentry", "Amplitude", "Playwright", "Vitest"}},
		{Name: "También", Items: alsoItems},
		{Name: "Idiomas", Items: []string{"Español (nativo)", "Francés (bilingüe)", "Inglés (profesional)"}},
	},
	i18n.FR: {
		{Name: "IA & dev agentique", Items: []string{"Claude Code", "Agents LLM (Claude, GPT/Codex)", "MCP", "Code review assisté par IA", "Conventions IA d'équipe (CLAUDE.md)"}},
		{Name: "Frontend", Items: frontendItems},
		{Name: "Backend & plateforme", Items: []string{".NET (C#)", "EF Core", "Node.js", "PostgreSQL", "API REST", "CI/CD (GitHub Actions)", "Vercel", "Sentry", "Amplitude", "Playwright", "Vitest"}},
		{Name: "Aussi", Items: alsoItems},
		{Name: "Langues", Items: []string{"Espagnol (natif)", "Français (bilingue)", "Anglais (professionnel)"}},
	},
}

var (
	React = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "react",
			Label:       "React",
			Alternative: "React Icon, an atom illustration in blue color",
			Source:      helpers.RepoURL + "/icons/react.svg",
		},
		Link: "https://react.dev/",
	}
	JavaScript = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "javaScript",
			Label:       "JavaScript",
			Alternative: "JavaScript Icon, white J S letters, yellow background",
			Source:      helpers.RepoURL + "/icons/javascript.svg",
		},
		Link: "https://es.wikipedia.org/wiki/JavaScript",
	}
	TypeScript = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "typeScript",
			Label:       "TypeScript",
			Alternative: "TypeScript Icon, white T S letters, blue background",
			Source:      helpers.RepoURL + "/icons/typescript.svg",
		},
		Link: "https://www.typescriptlang.org/",
	}
	Redux = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "redux",
			Label:       "Redux",
			Alternative: "Redux icon, a purple triquetra with dots",
			Source:      helpers.RepoURL + "/icons/redux.svg",
		},
		Link: "https://redux.js.org/",
	}
	TanStack = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "reactQuery",
			Label:       `TanStack Query`,
			Alternative: "TanStack query old icon that is as a flower with yellow center and red petals",
			Source:      helpers.RepoURL + "/icons/reactQuery.svg",
		},
		Link: "https://tanstack.com/query/latest/docs/framework/react/overview",
	}
	AntDesign = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "antD",
			Label:       "Ant Design",
			Alternative: "Ant design logo, a sided square with three quartes blue lines, the rest in red and a red dot at the center",
			Source:      "https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg",
		},
		Link: "https://ant.design/",
	}
	Node = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "node",
			Label:       "Node",
			Alternative: "Node logo, a green hexagon that links with the JavaScript logo",
			Source:      helpers.RepoURL + "/icons/node.svg",
		},
		Link: "https://nodejs.org/en/",
	}
	Go = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "go",
			Label:       "Go",
			Alternative: "Golang logo is the GO word in blue color",
			Source:      helpers.RepoURL + "/icons/golang.svg",
		},
		Link: "https://go.dev/",
	}
	MySQL = Skill{
		IconLabelParams: helpers.IconLabelParams{
			Id:          "mysql",
			Label:       "MySQL",
			Alternative: "MySQL logo, blue 'my' word and orange 'sql' in cappital letters with a dolphin at the top right corner",
			Source:      helpers.RepoURL + "/icons/mysql.svg",
		},
		Link: "https://www.mysql.com/",
	}
)
