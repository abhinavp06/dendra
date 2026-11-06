package main

import (
	"fmt"
	"net/http"
)

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", healthHandler)
	mux.HandleFunc("POST /claims/register", registerClaimHandler)

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func registerClaimHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "In claim handler")
}
