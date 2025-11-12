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

// CreateWalletHandler handles business wallet creation
func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	businessIDValue := r.Context().Value("businessID")
	if businessIDValue == nil {
		http.Error(w, "businessID not found in context", http.StatusUnauthorized)
		return
	}
	businessID := businessIDValue.(string)

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	conn, err := grpc.Dial("wallet-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to wallet-service: %v", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()

	client := pb.NewWalletServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.CreateWallet(ctx, &pb.CreateWalletRequest{
		BusinessId: businessID,
		Type:       pb.WalletType_BUSINESS,
	})
	if err != nil {
		log.Printf("error creating business wallet: %v", err)
		http.Error(w, "failed to create wallet", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// CreateCustomerWalletHandler handles customer wallet creation
func CreateCustomerWalletHandler(w http.ResponseWriter, r *http.Request) {
	businessIDValue := r.Context().Value("businessID")
	if businessIDValue == nil {
		http.Error(w, "businessID not found in context", http.StatusUnauthorized)
		return
	}
	businessID := businessIDValue.(string)

	customerID := r.URL.Path[len("/v0/wallet/"):]
	if customerID == "" {
		http.Error(w, "customerId required in path", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	conn, err := grpc.Dial("wallet-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to wallet-service: %v", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()

	client := pb.NewWalletServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.CreateWallet(ctx, &pb.CreateWalletRequest{
		BusinessId: businessID,
		CustomerId: customerID,
		Type:       pb.WalletType_CUSTOMER,
	})
	if err != nil {
		log.Printf("error creating customer wallet: %v", err)
		http.Error(w, "failed to create wallet", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
