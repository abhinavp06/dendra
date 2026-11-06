package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhinavp06/dendra/gateway/internal/config"
    "github.com/abhinavp06/dendra/gateway/internal/routing/base"
)

func main() {
	fmt.Println("starting dendra-gateway...")

	cfg := config.Load()

    http.HandleFunc("/", routing.BaseHandler)

    fmt.Printf("dendra-gateway listening on port %s\n", cfg.Port)

    addr := ":" + cfg.Port

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
