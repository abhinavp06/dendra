package main

import (
	"fmt"
	"net/http"
)

func registerClaimHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "In claim handler")
}
