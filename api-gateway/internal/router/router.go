package router

// import (
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/ronexlemon/rail/client-gateway/internal/handler"
// 	"github.com/ronexlemon/rail/client-gateway/internal/middleware"
// )

// func SetupRouter() *mux.Router {
// 	r := mux.NewRouter()

// 	// Apply middleware globally
// 	api := r.PathPrefix("/api").Subrouter()
// 	api.Use(middleware.AuthMiddleware)

// 	// Routes
// 	api.HandleFunc("/business", handler.GetBusinessDetails).Methods("GET")
// 	api.HandleFunc("/payments", handler.MakePayment).Methods("POST")

// 	return r
// }
