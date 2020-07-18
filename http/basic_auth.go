package http

import (
	"github.com/psolru/terrastate-http/config"
	"net/http"
)

// Set basic auth header and check if user is authenticated
func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)

		if isAuthorized(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

// Do the actual auth checks
func isAuthorized(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	// Check if any auth is received at all, if not - duh...
	if !ok {
		return false
	}

	// Check if username is correct
	if username != config.Values.Username {
		return false
	}

	// Check if password is correct
	if password != config.Values.Password {
		return false
	}
	return true
}
