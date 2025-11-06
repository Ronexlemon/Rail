# --- Build Stage ---
FROM golang:1.25-alpine AS builder


WORKDIR /app

RUN apk add --no-cache git
# Copy go.mod and go.sum for service + shared
COPY micro-services/business-service/go.mod ./micro-services/business-service/go.sum ./
COPY shared/go.mod ./shared/go.sum ./shared/

# Download dependencies
RUN go mod download

# Copy service source code and shared code
COPY micro-services/business-service/. ./
COPY shared ./shared

RUN go run github.com/steebchen/prisma-client-go generate
# Tidy dependencies and build binary
RUN go mod tidy
RUN go build -o business-service ./cmd/main.go

# Copy configs for runtime
COPY ./micro-services/business-service/configs ./configs

# --- Run Stage ---
FROM alpine:3.19

WORKDIR /app

# Copy binary and configs from builder
COPY --from=builder /app/business-service .
COPY --from=builder /app/configs ./configs

# Set environment variables
ENV SERVICE_NAME=business-service
ENV SERVICE_PORT=8085
ENV DATABASE_URL=postgres://admin:admin@postgres:5432/main?sslmode=disable
ENV KAFKA_BROKERS=kafka:9092

EXPOSE 8085

CMD ["./business-service"]
