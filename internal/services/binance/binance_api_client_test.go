package binance_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sergioazevedo/argentum_go/internal/services/binance"
	"github.com/stretchr/testify/assert"
)

const fakeTradeData = `[
	{
			"id": 136839513,
			"price": "53869.79000000",
			"qty": "0.00039000",
			"quoteQty": "21.00921810",
			"time": 1726338631723,
			"isBuyerMaker": true,
			"isBestMatch": true
	},
	{
			"id": 136839514,
			"price": "53878.36000000",
			"qty": "0.00086000",
			"quoteQty": "46.33538960",
			"time": 1726338638090,
			"isBuyerMaker": false,
			"isBestMatch": true
	}
]`

func fakeServer(statusCode int, responseData string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(responseData))
	}))

	return server
}

func TestAPIClient_FetchRecentTrades_Returns_Empty_List_Request_Fails(t *testing.T) {
	server := fakeServer(http.StatusBadRequest, "Error")

	client := binance.NewClient(server.URL)
	tradeList, err := client.FetchRecentTrades("BTCEUR", 2)

	server.Close()

	assert.Empty(t, tradeList)
	assert.NotNil(t, err)
}

func TestAPIClient_FetchRecentTrades_Returns_Trade_List_Request_OK(t *testing.T) {

	server := fakeServer(http.StatusOK, fakeTradeData)

	client := binance.NewClient(server.URL)
	tradeList, err := client.FetchRecentTrades("BTCEUR", 2)

	server.Close()

	assert.NotEmpty(t, tradeList)
	assert.Len(t, tradeList, 2)
	assert.Nil(t, err)
}
