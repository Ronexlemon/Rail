package main

import (
	"fmt"
	"net/http"
	//"strings"

	"github.com/ronexlemon/rail/api-gateway/internal/handler"
	"github.com/ronexlemon/rail/api-gateway/internal/middleware"
	"github.com/ronexlemon/rail/api-gateway/internal/router"
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

	// --- 2.3 Protected Wallet Routes ---

	// POST /v0/wallet (Create Business Wallet) | GET /v0/wallet (Get Business Wallets)
	mux.Handle("/v0/wallet", middleware.AuthMiddleware(http.HandlerFunc(router.WalletsRouteHandler)))
	
	// GET /v0/wallet/balance (Check Balances)
	// Must be defined before the generic /v0/wallet/ handler to take precedence
	mux.Handle("/v0/wallet/balance", middleware.AuthMiddleware(http.HandlerFunc(router.BalanceRouteHandler)))

	// POST/GET /v0/wallet/{customerId} (Customer Wallet Operations)
	// WARNING: ServeMux's path matching means this will catch EVERYTHING starting with /v0/wallet/
	// including /v0/wallet/balance if it wasn't defined first. This is a limitation.
	mux.Handle("/v0/wallet/", middleware.AuthMiddleware(http.HandlerFunc(router.CustomerWalletsRouteHandler)))

	fmt.Println("API Gateway running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
