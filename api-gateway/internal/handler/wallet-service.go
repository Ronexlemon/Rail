package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/ronexlemon/rail/micro-services/wallet-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// WalletRequest represents the JSON body for creating a wallet
type WalletRequest struct {
	BusinessId string `json:"business_id,omitempty"`
	CustomerId string `json:"customer_id,omitempty"`
	Type       string `json:"type"` // BUSINESS or CUSTOMER
}

// CreateWalletHandler creates a wallet via gRPC
func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON body
	var reqBody WalletRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Connect to gRPC Wallet Service
	conn, err := grpc.Dial("wallet-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to wallet-service: %v", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()

	client := pb.NewWalletServiceClient(conn)

	// Map type string to proto enum
	var walletType pb.WalletType
	switch reqBody.Type {
	case "BUSINESS":
		walletType = pb.WalletType_BUSINESS
	case "CUSTOMER":
		walletType = pb.WalletType_CUSTOMER
	default:
		http.Error(w, "invalid wallet type", http.StatusBadRequest)
		return
	}

	// Call gRPC method
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.CreateWallet(ctx, &pb.CreateWalletRequest{
		BusinessId: reqBody.BusinessId,
		CustomerId: reqBody.CustomerId,
		Type:       walletType,
	})
	if err != nil {
		log.Printf("error calling CreateWallet: %v", err)
		http.Error(w, "failed to create wallet", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
