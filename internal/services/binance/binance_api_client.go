package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sergioazevedo/argentum_go/internal/models"
)

const baseURL string = "https://data-api.binance.vision/api/v3"

type BinanceAPIClient struct{}

func (client *BinanceAPIClient) FetchRecentTrades(pair string, limit uint16) ([]models.BinanceTrade, error) {
	url := fmt.Sprintf(baseURL+"/trades?symbol=%s&limit=%d", pair, limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := performRequest(req)
	if err != nil {
		return nil, err
	}

	list := make([]models.BinanceTrade, 0, limit)
	json.NewDecoder(resp.Body).Decode(&list)

	return list, nil
}

func performRequest(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("%s", resp.Body)
	}

	return resp, nil
}
