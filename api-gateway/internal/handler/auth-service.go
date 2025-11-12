package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	proto "github.com/ronexlemon/rail/micro-services/auth-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// BusinessRequest is the incoming HTTP JSON body
type BusinessRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role    string `json:"role"`
	Password string `json:"pass"`
}

// RegisterBusinessHandler registers a new business via gRPC
func RegisterBusinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request
	var reqBody BusinessRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Connect to gRPC Auth Service
	conn, err := grpc.Dial("auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to auth-service: %v", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)

	// Call gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.RegisterUser(ctx, &proto.RegisterUserRequest{
		Email: reqBody.Email,
		Password: reqBody.Password,
		Name: reqBody.Name,
		Role: reqBody.Role,
		
	})
	if err != nil {
		log.Printf("error calling RegisterBusiness: %v", err)
		http.Error(w, "failed to register business", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
