package routing

import (
	"fmt"
	"net/http"
)

// HEALTH CHECK
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

// CLAIMS ROUTES
func RegisterClaimHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "In claim handler")
}
