package curriculumvitae

import "embed"

// Assets holds every static file (css, js, images, icons, CV PDFs, fonts)
// embedded into the binary so they are served by the deployment itself
// instead of being hotlinked from raw.githubusercontent.com/.../master.
// FS paths keep the "assets/" prefix (e.g. assets/css/output.css), which
// matches the "/assets/..." request paths, so no path rewriting is needed.
//
//go:embed all:assets
var Assets embed.FS
