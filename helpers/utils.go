package helpers

import "github.com/a-h/templ"

// RepoURL is the base path for static assets. Assets are served from the
// deployment itself (Echo static locally, embed.FS on Vercel) at /assets,
// so previews and any branch resolve correctly (no coupling to master).
var RepoURL templ.SafeURL = "/assets"
