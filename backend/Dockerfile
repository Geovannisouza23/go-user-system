# Build stage
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .
COPY .env .env

EXPOSE 8080

CMD ["./api"]
