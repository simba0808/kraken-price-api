package main

import (
	"log"
	"net/http"

	"github.com/simba0808/btc-ltp-service/internal/api"
	"github.com/simba0808/btc-ltp-service/internal/kraken"
)

func main() {
	krakenClient := kraken.NewClient()
	handler := api.NewHandler(krakenClient)

	http.HandleFunc("/api/v1/ltp", handler.GetLTP)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
