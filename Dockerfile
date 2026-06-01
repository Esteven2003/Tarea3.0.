# Etapa 1: Compilar la aplicación Go
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
# Deshabilitamos CGO para obtener un binario estático y más ligero
RUN CGO_ENABLED=0 GOOS=linux go build -o mi-accion main.go

# Etapa 2: Imagen mínima para ejecutar la acción
FROM alpine:latest
# GitHub Actions requiere ciertas herramientas básicas que Alpine ya trae
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/mi-accion /mi-accion

# El ENTRYPOINT le dice a Docker qué ejecutar cuando inicie el contenedor
ENTRYPOINT ["/mi-accion"]