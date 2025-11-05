# --- Build Stage ---
FROM golang:1.25-alpine AS builder


WORKDIR /app

# Copy go.mod and go.sum for service + shared
COPY micro-services/transaction-service/go.mod ./micro-services/transaction-service/go.sum ./
COPY shared/go.mod ./shared/go.sum ./shared/

# Download dependencies
RUN go mod download

# Copy service source code and shared code
COPY micro-services/transaction-service/. ./
COPY shared ./shared

# Ensure dependencies are tidy and build binary
RUN go mod tidy
RUN go build -o transaction-service ./cmd/main.go

# Copy configs for runtime
COPY ./micro-services/transaction-service/configs ./configs

# --- Run Stage ---
FROM alpine:3.19

WORKDIR /app

# Copy binary and configs from builder
COPY --from=builder /app/transaction-service .
COPY --from=builder /app/configs ./configs

# Set environment variables
ENV SERVICE_NAME=transaction-service
ENV SERVICE_PORT=8082
ENV DATABASE_URL=postgres://admin:admin@postgres:5432/main?sslmode=disable
ENV KAFKA_BROKERS=kafka:9092

EXPOSE 8082

CMD ["./transaction-service"]
