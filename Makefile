.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

.PHONY: tailwind-minify
tailwind-minify:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

.PHONY: dev-css
dev-css:
	make tailwind-minify && make tailwind-build

.PHONY: dev
dev:
	go build -o ./tmp/app.exe ./cmd/main.go && air

.PHONY: build
build:
	go mod tidy && make templ-generate && go build -o ./tmp/app.exe ./cmd/main.go
