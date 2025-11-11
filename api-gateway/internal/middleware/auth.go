package middleware

import (
	"context"

	"log"
	"net/http"
	"time"

	proto "github.com/ronexlemon/rail/micro-services/auth-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AuthMiddleware returns an http.Handler that checks apiKey & secretKey via gRPC
func AuthMiddleware(next http.Handler) http.Handler {
	// Create gRPC client connection (could be global for efficiency)
	conn, err := grpc.Dial("auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to auth-service: %v", err)
	}
	client := proto.NewAuthServiceClient(conn)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		secretKey := r.Header.Get("X-Secret-Key")

		if apiKey == "" || secretKey == "" {
			http.Error(w, "missing API keys", http.StatusUnauthorized)
			return
		}

		// Call gRPC auth-service
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		resp, err := client.GetBusinessByKeys(ctx, &proto.GetBusinessByKeysRequest{
			ApiKey:    apiKey,
			SecretKey: secretKey,
		})
		if err != nil || resp.Status != "ACTIVE" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Optionally attach business ID to context for downstream handlers
		ctx = context.WithValue(r.Context(), "businessID", resp.BusinessId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
