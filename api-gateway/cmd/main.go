package main

import (
	"net/http"

	"github.com/ronexlemon/rail/api-gateway/internal/middleware"
)


func main() {
	mux := http.NewServeMux()

	// Protected route
	mux.Handle("/business/data", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		businessID := r.Context().Value("businessID").(string)
		w.Write([]byte("Hello Business: " + businessID))
	})))

	http.ListenAndServe(":8080", mux)
	
}