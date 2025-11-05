# --- Build Stage ---
FROM golang:1.25-alpine AS builder


WORKDIR /app

# Copy go.mod and go.sum for service + shared
COPY micro-services/settlement-service/go.mod ./micro-services/settlement-service/go.sum ./
COPY shared/go.mod ./shared/go.sum ./shared/

# Download dependencies
RUN go mod download

# Copy service source code and shared code
COPY micro-services/settlement-service/. ./
COPY shared ./shared

# Tidy dependencies and build binary
RUN go mod tidy
RUN go build -o settlement-service ./cmd/main.go

# Copy configs for runtime
COPY ./micro-services/settlement-service/configs ./configs

# --- Run Stage ---
FROM alpine:3.19

WORKDIR /app

# Copy binary and configs from builder
COPY --from=builder /app/settlement-service .
COPY --from=builder /app/configs ./configs

# Set environment variables
ENV SERVICE_NAME=settlement-service
ENV SERVICE_PORT=8084
ENV DATABASE_URL=postgres://admin:admin@postgres:5432/main?sslmode=disable
ENV KAFKA_BROKERS=kafka:9092

EXPOSE 8084

CMD ["./settlement-service"]
