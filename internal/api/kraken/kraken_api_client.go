package kraken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/sergioazevedo/argentum_go/internal/lib/http_request"
)

const DEFAULT_BASE_URL string = "https://api.kraken.com/0/public"

type APIClient struct {
	BaseURL string
}

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error  []string
	Result map[string]([][]interface{})
}

type Trade struct {
	Date     time.Time
	Volume   decimal.Decimal
	Price    decimal.Decimal
	Quantity decimal.Decimal
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

func (c APIClient) FetchRecentTrades(pair string, limit int16) ([]Trade, error) {
	url := fmt.Sprintf(c.BaseURL+"/Trades?pair=%s&count=%d", pair, limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http_request.Perform(req)
	if err != nil {
		return nil, err
	}

	jsonData := KrakenResponse{}
	json.NewDecoder(resp.Body).Decode(&jsonData)
	trades := jsonData.Result[pair]

	result := make([]Trade, 0, len(trades))
	for _, trade := range trades {
		price, err := decimal.NewFromString(trade[0].(string))
		if err != nil {
			return nil, err
		}

		volume, err := decimal.NewFromString(trade[1].(string))
		if err != nil {
			return nil, err
		}

		date := strconv.FormatFloat(trade[2].(float64), 'f', -1, 64)
		dateParts := strings.Split(date, ".")
		seconds, err := strconv.ParseInt(dateParts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		milliseconds, err := strconv.ParseInt(dateParts[1], 10, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, Trade{
			Date:     time.Unix(seconds, milliseconds).UTC(),
			Volume:   volume,
			Price:    price,
			Quantity: volume.Div(price),
		})
	}

	return result, nil
}
