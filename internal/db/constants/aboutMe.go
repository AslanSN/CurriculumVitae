package constants

import "github.com/AslanSN/CurriculumVitae/internal/i18n"

// AboutMe is the "About" paragraph, per locale. ES/FR are drafted from the EN
// source and the trilingual CV PDFs — review the copy before publishing.
var AboutMe = map[i18n.Locale]string{
	i18n.EN: `Senior full-stack engineer (~4.5 years) who ships end to end — React 19 / Next.js + TypeScript on the front, real .NET / PostgreSQL on the back, plus platform work (CI/CD, testing, observability). I work AI-native: Claude Code, LLM agents and MCP are part of how I ship, and I authored my team's shared AI-engineering conventions. At a smart-building SaaS I owned frontend delivery and sustained it solo for ~8 months, and in my final stretch took over cross-department delivery coordination previously held by the CTO. Trilingual (Spanish native, French bilingual, English professional), remote-first, with a humanities background I bring to communication and product judgment.`,

	i18n.ES: `Ingeniero full-stack senior (~4,5 años) que entrega de punta a punta — React 19 / Next.js + TypeScript en el front, .NET / PostgreSQL real en el back, más trabajo de plataforma (CI/CD, testing, observabilidad). Trabajo AI-native: Claude Code, agentes LLM y MCP forman parte de cómo entrego, y redacté las convenciones de ingeniería con IA de mi equipo. En un SaaS de edificios inteligentes lideré la entrega del frontend y la sostuve en solitario durante ~8 meses, y en mi tramo final asumí la coordinación de entrega entre departamentos que antes llevaba el CTO. Trilingüe (español nativo, francés bilingüe, inglés profesional), remote-first, con una formación en humanidades que aplico a la comunicación y al criterio de producto.`,

	i18n.FR: `Ingénieur full-stack senior (~4,5 ans) qui livre de bout en bout — React 19 / Next.js + TypeScript côté front, un vrai backend .NET / PostgreSQL, ainsi que du travail de plateforme (CI/CD, tests, observabilité). Je travaille AI-native : Claude Code, agents LLM et MCP font partie de ma façon de livrer, et j'ai rédigé les conventions d'ingénierie IA de mon équipe. Dans un SaaS de bâtiments intelligents, j'ai piloté la livraison du frontend et l'ai maintenue seul pendant ~8 mois, et dans ma dernière période j'ai repris la coordination des livraisons entre départements auparavant assurée par le CTO. Trilingue (espagnol natif, français bilingue, anglais professionnel), remote-first, avec une formation en sciences humaines que j'applique à la communication et au jugement produit.`,
}

// AslanImageAlt is the alt text for the hero portrait, per locale.
var AslanImageAlt = map[i18n.Locale]string{
	i18n.EN: "Close-up photo of Alan Staub Negro smiling",
	i18n.ES: "Foto en primer plano de Alan Staub Negro sonriendo",
	i18n.FR: "Photo en gros plan d'Alan Staub Negro souriant",
}
