.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-generate-watch
templ-generate-watch:
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

# .PHONY: build
# build:
#	go mod tidy && make templ-generate-watch && make dev-css && make go-build

.PHONY: build
build:
	go mod tidy && make templ-generate && make go-build && make dev-css
