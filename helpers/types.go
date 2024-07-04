package helpers

import "github.com/a-h/templ"

type IconLabelParams struct {
	Id, Label, Alternative string
	Source                 templ.SafeURL
}
