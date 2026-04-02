package handlers

import (
	"net/http"
	"os"
)

// WithCORS wraps h: sets CORS headers, answers OPTIONS with 204, then forwards other methods.
func WithCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
		if allowOrigin == "" {
			allowOrigin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if reqHdrs := r.Header.Get("Access-Control-Request-Headers"); reqHdrs != "" {
			w.Header().Set("Access-Control-Allow-Headers", reqHdrs)
		} else {
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		}
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
