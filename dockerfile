# syntax=docker/dockerfile:1

# Etapa de construcción
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instalar make, templ y cualquier otra dependencia necesaria
RUN apk update && apk add --no-cache make gcc libc-dev git

# Instalar templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copiar y descargar dependencias
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar el resto de los archivos
COPY . .

# Ejecutar Makefile para construir la aplicación
RUN make build

# Etapa final para la imagen
FROM golang:1.22-alpine

WORKDIR /app

# Copiar binario desde la etapa de construcción
COPY --from=builder /app/tmp/app.exe /app/app.exe

# Configurar Air para desarrollo
RUN go install github.com/air-verse/air@latest
# RUN go mod tidy
COPY .air.toml /app/.air.toml

EXPOSE 2340

# Comando para producción
# CMD [ "/app" ]

# Comando para desarrollo (opcional, descomentar si quieres usar air)
CMD [ "air" ]
