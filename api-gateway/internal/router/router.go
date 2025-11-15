package router



import (
	
	"net/http"
	"strings"

	"github.com/ronexlemon/rail/api-gateway/internal/handler"
	
)

// --- 1. Dedicated Handler Wrappers for Complex Routes ---

// WalletsRouteHandler handles requests to /v0/wallet based on the HTTP method.
func WalletsRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// POST /v0/wallet -> Create business wallet
		handler.CreateWalletHandler(w, r)
	case http.MethodGet:
		// GET /v0/wallet -> Fetch all business wallets
		handler.WalletsHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// CustomerWalletsRouteHandler handles requests to /v0/wallet/{customerId} and extracts the customerID.
func CustomerWalletsRouteHandler(w http.ResponseWriter, r *http.Request) {
	const prefix = "/v0/wallet/"
	customerID := strings.TrimPrefix(r.URL.Path, prefix)
	
	if customerID == "" || customerID == "balance" {
		// If the path is exactly /v0/wallet/ or /v0/wallet/balance, it's not a customer ID route
		// If you only expect this handler to run on /v0/wallet/{id}, you need a better router.
		// For ServeMux, this is a limitation, but we can try to redirect or error.
		http.Error(w, "invalid or missing customerId in path", http.StatusBadRequest)
		return
	}
	
	// Inject customerID into the request context for the handler functions to use
	// NOTE: handler.CreateCustomerWalletHandler and handler.CustomerWalletsHandler 
	// must be updated to read the customerID from the path instead of the context.
	
	switch r.Method {
	case http.MethodPost:
		// POST /v0/wallet/{customerId} -> Create customer wallet
		// You may want this to be POST /v0/customer/wallet or PUT /v0/wallet/{customerId}
		handler.CreateCustomerWalletHandler(w, r)
	case http.MethodGet:
		// GET /v0/wallet/{customerId} -> Fetch customer wallets
		handler.CustomerWalletsHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// BalanceRouteHandler handles requests to /v0/wallet/balance.
func BalanceRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// GET /v0/wallet/balance -> Fetch wallet balances
		// Optional query parameters: ?customerId=xxx&network=evm
		handler.WalletsChainBalanceHandler(w, r)
	default:
		http.Error(w, "method not allowedd", http.StatusMethodNotAllowed)
	}
}
