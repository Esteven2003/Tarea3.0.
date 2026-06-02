# Etapa 1: Compilar el código de Go
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o calculadora .

# Etapa 2: Crear la imagen final ligera
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/calculadora .
EXPOSE 8080
CMD ["./calculadora"]