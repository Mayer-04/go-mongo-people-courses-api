package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mayer-04/go-mongo-people-courses-api/internal/config"
)

func APIKey(next http.HandlerFunc, cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")

		if key == "" || key != cfg.ApiKey {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)

			if err := json.NewEncoder(w).Encode(APIKeyErrorResponse{
				Error:   http.StatusText(http.StatusForbidden),
				Message: "invalid or missing API key",
			}); err != nil {
				log.Println("error writing JSON:", err)
			}

			return
		}

		next.ServeHTTP(w, r)
	}
}
