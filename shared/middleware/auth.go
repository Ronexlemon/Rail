package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type BusinessInfo struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// define a private key type to avoid collisions
type contextKey string

// define a constant for the business key
const businessCtxKey contextKey = "business"

func ValidateBusinessFromService(businessServiceURL string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			secretKey := r.Header.Get("X-Secret-Key")

			if apiKey == "" || secretKey == "" {
				http.Error(w, "Missing API credentials", http.StatusUnauthorized)
				return
			}

			client := &http.Client{Timeout: 5 * time.Second}
			req, _ := http.NewRequest("GET", businessServiceURL+"/internal/validate", nil)
			req.Header.Set("X-API-Key", apiKey)
			req.Header.Set("X-Secret-Key", secretKey)

			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				http.Error(w, "Invalid or inactive business credentials", http.StatusUnauthorized)
				return
			}
			defer resp.Body.Close()

			var business BusinessInfo
			json.NewDecoder(resp.Body).Decode(&business)

			if business.Status != "ACTIVE" {
				http.Error(w, "Business is not active", http.StatusForbidden)
				return
			}

			
			ctx := context.WithValue(r.Context(), businessCtxKey, business)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// helper to retrieve business info from context safely
func GetBusinessFromContext(ctx context.Context) (*BusinessInfo, bool) {
	business, ok := ctx.Value(businessCtxKey).(BusinessInfo)
	return &business, ok
}
