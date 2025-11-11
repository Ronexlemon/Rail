package middleware

import (
	"context"
	"net/http"
	"time"

	pb "github.com/ronexlemon/rail/micro-services/auth-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BusinessInfoKey string

type BusinessInfo struct {
	BusinessID string
	Status     string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		secretKey := r.Header.Get("X-Secret-Key")
		if apiKey == "" || secretKey == "" {
			http.Error(w, "Missing API key or Secret key", http.StatusUnauthorized)
			return
		}

		// Use context with timeout for the connection
		connCtx, cancelConn := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelConn()

		conn, err := grpc.DialContext(
			connCtx,
			"auth-service:50052",
			grpc.WithTransportCredentials(insecure.NewCredentials()), // non-deprecated
			grpc.WithBlock(),                                         // wait until connection ready
		)
		if err != nil {
			http.Error(w, "Failed to connect to auth service", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		client := pb.NewAuthServiceClient(conn)

		// Call Authenticate with a timeout context
		callCtx, cancelCall := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancelCall()

		resp, err := client.Authenticate(callCtx, &pb.AuthenticateRequest{
			ApiKey:    apiKey,
			SecretKey: secretKey,
		})
		if err != nil || !resp.Valid {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Store business info in request context
		bizInfo := &BusinessInfo{
			BusinessID: resp.BusinessId,
			Status:     resp.Status,
		}

		ctx := context.WithValue(r.Context(), BusinessInfoKey("businessInfo"), bizInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
