TW := ./bin/tailwindcss
CSS_IN := ./assets/css/input.css
CSS_OUT := ./assets/css/output.css
TEMPL := $(shell go env GOPATH)/bin/templ

.PHONY: tw-install templ-generate css css-watch build build-prod dev

## download the Tailwind v4 standalone CLI (no Node needed)
tw-install:
	mkdir -p bin && curl -sLo $(TW) https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64 && chmod +x $(TW)

templ-generate:
	$(TEMPL) generate

## production stylesheet — commit the result (Vercel @vercel/go has no build hook)
css:
	$(TW) -i $(CSS_IN) -o $(CSS_OUT) --minify

## run in a second terminal during development
css-watch:
	$(TW) -i $(CSS_IN) -o $(CSS_OUT) --watch

## fast build for air (CSS handled separately by css-watch)
build:
	$(TEMPL) generate && go build -o ./tmp/app.exe ./cmd/main.go

## full build before committing / deploying
build-prod:
	go mod tidy && $(TEMPL) generate && $(MAKE) css && go build -o ./tmp/app.exe ./cmd/main.go

dev:
	air
