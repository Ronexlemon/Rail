# --- Build Stage ---
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Install git for dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum for API Gateway + shared
COPY api-gateway/go.mod ./api-gateway/go.sum ./
COPY shared/go.mod ./shared/go.sum ./shared/

# Download dependencies
RUN go mod download

# Copy source code
COPY api-gateway/. ./
COPY shared ./shared

# Tidy modules and build binary
RUN go mod tidy
RUN go build -o api-gateway ./cmd/main.go

# Copy any config files if needed
COPY api-gateway/configs ./configs

# --- Run Stage ---
FROM alpine:3.19

WORKDIR /app

# Copy binary and configs from builder
COPY --from=builder /app/api-gateway .
COPY --from=builder /app/configs ./configs

# Set environment variables
ENV SERVICE_NAME=api-gateway
ENV SERVICE_PORT=8080
ENV AUTH_GRPC=auth-service:50051
ENV BUSINESS_GRPC=business-service:50052
ENV TRANSACTION_GRPC=transaction-service:50053

# Expose HTTP port
EXPOSE 8080

# Run the binary
CMD ["./api-gateway"]
