.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ
templ:
	templ generate --watch --proxy=http://localhost:2345

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

.PHONY: tailwind-minify
tailwind-minify:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

.PHONY: dev-css
dev-css:
	make tailwind-minify && make tailwind-build

.PHONY: go-build
go-build:
	go build -o ./tmp/app.exe ./cmd/main.go

.PHONY: dev
dev:
	make go-build && air
	

.PHONY: build
build:
	go mod tidy && make templ-generate
