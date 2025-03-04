package kraken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

type krakenResponse struct {
	Error  []string               `json:"error"`
	Result map[string]interface{} `json:"result"`
}

func (c *Client) GetLTP(pair string) (float64, error) {
	krakenPair := strings.Replace(pair, "/", "%2f", -1)
	url := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%s", krakenPair)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var krakenResp krakenResponse
	if err := json.NewDecoder(resp.Body).Decode(&krakenResp); err != nil {
		return 0, err
	}

	if len(krakenResp.Error) > 0 {
		return 0, fmt.Errorf("Kraken API error: %s", strings.Join(krakenResp.Error, ", "))
	}

	pairData, ok := krakenResp.Result[pair]
	if !ok {
		return 0, fmt.Errorf("no data for pair %s", pair)
	}

	pairDataMap, ok := pairData.(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid data format for pair %s", pair)
	}

	lastTrade, ok := pairDataMap["c"]
	if !ok {
		return 0, fmt.Errorf("no last trade price for pair %s", pair)
	}

	lastTradeSlice, ok := lastTrade.([]interface{})
	if !ok || len(lastTradeSlice) < 1 {
		return 0, fmt.Errorf("invalid last trade price format for pair %s", pair)
	}

	lastTradeStr, ok := lastTradeSlice[0].(string)
	if !ok {
		return 0, fmt.Errorf("invalid last trade price type for pair %s", pair)
	}

	price, err := strconv.ParseFloat(lastTradeStr, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse price for pair %s: %v", pair, err)
	}

	return price, nil
}
