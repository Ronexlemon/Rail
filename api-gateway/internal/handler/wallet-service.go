package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

//Wallets

func WalletsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract businessID from context
	businessIDValue := r.Context().Value("businessID")
	if businessIDValue == nil {
		http.Error(w, "businessID not found in context", http.StatusUnauthorized)
		return
	}
	businessID := businessIDValue.(string)

	// Connect to Wallet Service via gRPC
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

	// Call BusinessWallet RPC
	resp, err := client.BusinessWallet(ctx, &pb.BusinessWalletRequest{
		BusinessId: businessID,
	})
	if err != nil {
		log.Printf("error fetching business wallets: %v", err)
		http.Error(w, "failed to fetch wallets", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


func CustomerWalletsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract businessID from context
	businessIDValue := r.Context().Value("businessID")
	if businessIDValue == nil {
		http.Error(w, "businessID not found in context", http.StatusUnauthorized)
		return
	}
	businessID := businessIDValue.(string)

	// Extract customerId from URL: /v0/wallet/{customerId}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[3] == "" {
		http.Error(w, "customerId required in path", http.StatusBadRequest)
		return
	}
	customerID := pathParts[3]

	// Connect to Wallet Service via gRPC
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

	// Call BusinessWallet RPC with customerId
	resp, err := client.BusinessWallet(ctx, &pb.BusinessWalletRequest{
		BusinessId: businessID,
		CustomerId: customerID,
	})
	if err != nil {
		log.Printf("error fetching customer wallets: %v", err)
		http.Error(w, "failed to fetch wallets", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}



//wallet balance

func WalletsBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract businessID from context
	businessIDValue := r.Context().Value("businessID")
	if businessIDValue == nil {
		http.Error(w, "businessID not found in context", http.StatusUnauthorized)
		return
	}
	businessID, ok := businessIDValue.(string)
	if !ok || businessID == "" {
		http.Error(w, "invalid businessID in context", http.StatusUnauthorized)
		return
	}

	// Connect to Wallet Service via gRPC
	conn, err := grpc.Dial("wallet-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to wallet-service: %v", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()

	client := pb.NewWalletServiceClient(conn)

	// Call WalletBalance RPC
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := client.WalletBalance(ctx, &pb.WalletBalanceRequest{
		BusinessId: businessID,
		Network:    "evm", // specify the network
	})
	if err != nil {
		log.Printf("error fetching wallet balances: %v", err)
		http.Error(w, "failed to fetch wallet balances", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

type RequestBody struct{
	Chain string `json:"chain"`
}


func WalletsChainBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract businessID from context
	businessIDValue := r.Context().Value("businessID")
	if businessIDValue == nil {
		http.Error(w, "businessID not found in context", http.StatusUnauthorized)
		return
	}
	businessID, ok := businessIDValue.(string)
	if !ok || businessID == "" {
		http.Error(w, "invalid businessID in context", http.StatusUnauthorized)
		return
	}

	var req RequestBody 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
 if req.Chain == ""{
	http.Error(w, "Missing chain", http.StatusBadRequest)
		return

 }
	// Connect to Wallet Service via gRPC
	conn, err := grpc.Dial("wallet-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to wallet-service: %v", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer conn.Close()

	client := pb.NewWalletServiceClient(conn)

	// Call WalletBalance RPC
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := client.WalletChainBalance(ctx, &pb.WalletBalanceRequest{
		BusinessId: businessID,
		Network:    "evm", // specify the network
		Chain: req.Chain,
	})
	if err != nil {
		log.Printf("error fetching wallet balances: %v", err)
		http.Error(w, "failed to fetch wallet balances", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
