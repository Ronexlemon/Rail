package middleware

import (
	"context"
	"net/http"
	"time"
	"log"

	businesspb "github.com/ronexlemon/rail/micro-services/business-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BusinessInfoKey string

type BusinessInfo struct {
	BusinessID string
	Status     string
}

// MiddlewareManager holds the persistent gRPC client
type MiddlewareManager struct {
	businessClient businesspb.BusinessServiceClient
}

// NewMiddlewareManager creates a persistent gRPC client connection
func NewMiddlewareManager(businessAddr string) *MiddlewareManager {
	conn, err := grpc.Dial(
		businessAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second), // initial dial timeout
	)
	if err != nil {
		log.Fatalf("Failed to connect to business-service: %v", err)
	}

	client := businesspb.NewBusinessServiceClient(conn)

	return &MiddlewareManager{
		businessClient: client,
	}
}

// AuthMiddleware returns a middleware using the persistent gRPC client
func (m *MiddlewareManager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		secretKey := r.Header.Get("X-Secret-Key")
		if apiKey == "" || secretKey == "" {
			http.Error(w, "Missing API key or Secret key", http.StatusUnauthorized)
			return
		}

		// Use a short timeout for each request
		callCtx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		resp, err := m.businessClient.GetBusinessByKeys(callCtx, &businesspb.GetBusinessByKeysRequest{
			ApiKey:    apiKey,
			SecretKey: secretKey,
		})
		if err != nil {
			http.Error(w, "Invalid API key or secret key", http.StatusUnauthorized)
			return
		}

		// Store business info in context
		bizInfo := &BusinessInfo{
			BusinessID: resp.BusinessId,
			Status:     resp.Status,
		}

		ctx := context.WithValue(r.Context(), BusinessInfoKey("businessInfo"), bizInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
