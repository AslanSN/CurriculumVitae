package i18n

import (
	"reflect"
	"testing"
)

// TestDictsComplete guarantees the type-safe promise: every supported locale
// has a fully-populated catalog with no empty strings. Adding a Dict field
// without translating it in all three languages fails here.
func TestDictsComplete(t *testing.T) {
	for _, loc := range Supported {
		d, ok := dicts[loc]
		if !ok {
			t.Fatalf("locale %q has no dictionary", loc)
		}
		v := reflect.ValueOf(d)
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).String() == "" {
				t.Errorf("locale %q: field %q is empty", loc, v.Type().Field(i).Name)
			}
		}
	}
}

func TestDetectAcceptLanguage(t *testing.T) {
	cases := []struct {
		header string
		want   Locale
	}{
		{"", EN},
		{"en-US,en;q=0.9", EN},
		{"es-ES,es;q=0.9,en;q=0.8", ES},
		{"fr-FR,fr;q=0.9", FR},
		{"de-DE,de;q=0.9", EN},              // unsupported → default
		{"en;q=0.7,es;q=0.9", ES},           // honors q ordering, not position
		{"pt-BR,pt;q=0.9,fr;q=0.4", FR},     // first supported wins
	}
	for _, c := range cases {
		if got := detectAcceptLanguage(c.header); got != c.want {
			t.Errorf("detectAcceptLanguage(%q) = %q, want %q", c.header, got, c.want)
		}
	}
}
