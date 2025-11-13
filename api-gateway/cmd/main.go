package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ronexlemon/rail/api-gateway/internal/handler"
	"github.com/ronexlemon/rail/api-gateway/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Protected example route
	mux.Handle("/business/data", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		businessID := r.Context().Value("businessID").(string)
		w.Write([]byte("Hello Business: " + businessID))
	})))

	// Public route to register a business
	mux.HandleFunc("/register-business", handler.RegisterBusinessHandler)

	// Wallet endpoints (protected)
	mux.Handle("/v0/wallet", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Create business wallet
			handler.CreateWalletHandler(w, r)
		} else if r.Method == http.MethodGet {
			// Fetch all business wallets
			handler.WalletsHandler(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Customer wallet endpoint: /v0/wallet/{customerId}
	mux.Handle("/v0/wallet/", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customerID := strings.TrimPrefix(r.URL.Path, "/v0/wallet/")
		if customerID == "" {
			http.Error(w, "customerId required in path", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			// Create customer wallet
			handler.CreateCustomerWalletHandler(w, r)
		} else if r.Method == http.MethodGet {
			// Fetch customer wallets
			handler.CustomerWalletsHandler(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))


	//balances
	mux.Handle("/v0/wallet/balance", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Create wallet (business or customer)
			//handler.CreateWalletHandler(w, r)
		case http.MethodGet:
			// Fetch wallet balances
			// Optional query parameters: ?customerId=xxx&network=evm
			handler.WalletsBalanceHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	fmt.Println("API Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
