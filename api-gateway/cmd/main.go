package main

import (
	"fmt"
	"net/http"

	"github.com/ronexlemon/rail/api-gateway/internal/handler"
	"github.com/ronexlemon/rail/api-gateway/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Protected route example
	mux.Handle("/business/data", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		businessID := r.Context().Value("businessID").(string)
		w.Write([]byte("Hello Business: " + businessID))
	})))

	// Register business
	mux.HandleFunc("/register-business", handler.RegisterBusinessHandler)

	// Wallet endpoints (protected)
	mux.Handle("/v0/wallet", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateWalletHandler)))

	mux.Handle("/v0/wallet/", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateCustomerWalletHandler)))

	fmt.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", mux)
}
