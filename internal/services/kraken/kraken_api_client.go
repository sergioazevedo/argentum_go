package kraken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sergioazevedo/argentum_go/internal/lib/http_request"
	"github.com/shopspring/decimal"
)

const DEFAULT_BASE_URL string = "https://api.kraken.com/0/public"

type APIClient struct {
	BaseURL string
}

type Trades [][]interface{}

func NewClient(baseUrl string) *APIClient {
	value := baseUrl
	if baseUrl == "" {
		value = DEFAULT_BASE_URL
	}

	return &APIClient{
		BaseURL: value,
	}
}

func (c APIClient) FetchRecentTrades(pair string, limit int16) (Trades, error) {
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

	return trades, nil
}

// KrakenResponse wraps the Kraken API JSON response
type KrakenResponse struct {
	Error  []string
	Result map[string]([][]interface{})
}

type Trade struct {
	date   time.Time
	volume decimal.Decimal
	price  decimal.Decimal
}

func (t Trade) Date() time.Time {
	return t.date
}

func (t Trade) Quantity() decimal.Decimal {
	return t.volume.Div(t.price)
}

func (t Trade) Price() decimal.Decimal {
	return t.price
}

func (t Trade) Volume() decimal.Decimal {
	return t.volume
}

func (t *Trade) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	_price := v[0].(string)
	_volume := v[1].(string)
	_date := int64(v[2].(float64))

	t.date = time.UnixMilli(_date)
	t.price, _ = decimal.NewFromString(_price)
	t.volume, _ = decimal.NewFromString(_volume)

	return nil
}
