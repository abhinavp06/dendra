package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting dendra-gateway...")

	cfg := loadConfig()

	fmt.Printf("dendra-gateway listening on port %s\n", cfg.Port)

	addr := ":" + cfg.Port

	if err := http.ListenAndServe(addr, newRouter()); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
