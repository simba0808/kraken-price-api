package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/simba0808/btc-ltp-service/internal/kraken"
)

type Handler struct {
	client *kraken.Client
}

func NewHandler(client *kraken.Client) *Handler {
	return &Handler{client: client}
}

type LTPResponse struct {
	LTP []struct {
		Pair  string  `json:"pair"`
		Price float64 `json:"price"`
	} `json:"ltp"`
}

func (h *Handler) GetLTP(w http.ResponseWriter, r *http.Request) {
	pairs := []string{"BTC/USD", "BTC/CHF", "BTC/EUR"}
	results := make([]struct {
		Pair  string  `json:"pair"`
		Price float64 `json:"price"`
	}, len(pairs))

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, pair := range pairs {
		wg.Add(1)
		go func(i int, pair string) {
			fmt.Println(pair)
			defer wg.Done()
			price, err := h.client.GetLTP(pair)
			if err == nil {
				mu.Lock()
				results[i] = struct {
					Pair  string  `json:"pair"`
					Price float64 `json:"price"`
				}{Pair: pair, Price: price}
				mu.Unlock()
			}
		}(i, pair)
	}

	wg.Wait()

	response := LTPResponse{LTP: results}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
