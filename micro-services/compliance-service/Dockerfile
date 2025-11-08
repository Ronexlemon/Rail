# --- Build Stage ---
FROM golang:1.25-alpine AS builder


WORKDIR /app

RUN apk add --no-cache git
# Copy go.mod and go.sum for service + shared
COPY micro-services/wallet-service/go.mod ./micro-services/wallet-service/go.sum ./
COPY shared/go.mod ./shared/go.sum ./shared/

# Download dependencies
RUN go mod download

# Copy service source and shared code
COPY micro-services/wallet-service/. ./
COPY shared ./shared


RUN go run github.com/steebchen/prisma-client-go generate
# Tidy dependencies and build
RUN go mod tidy
RUN go build -o wallet-service ./cmd/main.go

# Copy configs for runtime
COPY ./micro-services/wallet-service/configs ./configs

# --- Run Stage ---
FROM alpine:3.19

WORKDIR /app

# Copy binary and configs from builder
COPY --from=builder /app/wallet-service .
COPY --from=builder /app/configs ./configs

# Set environment variables
ENV SERVICE_NAME=wallet-service
ENV SERVICE_PORT=8083
ENV DATABASE_URL=postgres://admin:admin@postgres:5432/main?sslmode=disable
ENV KAFKA_BROKERS=kafka:9092

EXPOSE 8083

CMD ["./wallet-service"]
