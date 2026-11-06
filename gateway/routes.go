package main

import "net/http"

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", healthHandler)
	mux.HandleFunc("POST /claims/register", registerClaimHandler)

	return mux
}
