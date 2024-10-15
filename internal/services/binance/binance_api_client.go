package binance

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sergioazevedo/argentum_go/internal/lib/http_request"
)

const DEFAULT_BASE_URL string = "https://data-api.binance.vision/api/v3"

type APIClient struct {
	BaseURL string
}

type Trade struct {
	Price string  `json:"price"`
	Qty   string  `json:"qty"`
	Time  float64 `json:"time"`
}

func NewClient(baseUrl string) *APIClient {
	value := baseUrl
	if baseUrl == "" {
		value = DEFAULT_BASE_URL
	}

	return &APIClient{
		BaseURL: value,
	}
}

func (client *APIClient) FetchRecentTrades(pair string, limit uint16) ([]Trade, error) {
	url := fmt.Sprintf(client.BaseURL+"/trades?symbol=%s&limit=%d", pair, limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http_request.Perform(req)
	if err != nil {
		return nil, err
	}

	list := make([]Trade, 0, limit)
	json.NewDecoder(resp.Body).Decode(&list)

	return list, nil
}
