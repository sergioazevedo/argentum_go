package kraken_test

import (
	"net/http"
	"testing"

	"github.com/sergioazevedo/argentum_go/internal/lib/servertest"
	"github.com/sergioazevedo/argentum_go/internal/services/kraken"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const fakeTradeData = `{
	"error": [],
	"result": {
			"XXBTZEUR": [
					[
							"53960.00000",
							"0.00136938",
							1726341504.0383348,
							"b",
							"l",
							"",
							89154536
					],
					[
							"53959.90000",
							"0.01328528",
							1726341510.5618985,
							"s",
							"m",
							"",
							89154537
					]
			],
			"last": "1726349103032507974"
	}
}`

func TestAPIClient_FetchRecentTrades_Returns_Empty_List_Request_Fails(t *testing.T) {
	server := servertest.NewTestServer(http.StatusBadRequest, "Error")

	client := kraken.NewClient(server.URL)
	tradeList, err := client.FetchRecentTrades("XXBTZEUR", 2)

	server.Close()

	assert.Empty(t, tradeList)
	assert.NotNil(t, err)
}

func TestAPIClient_FetchRecentTrades_Returns_Trade_List_Request_OK(t *testing.T) {
	server := servertest.NewTestServer(http.StatusOK, fakeTradeData)

	client := kraken.NewClient(server.URL)
	tradeList, err := client.FetchRecentTrades("XXBTZEUR", 2)

	server.Close()

	assert.NotEmpty(t, tradeList)
	assert.Len(t, tradeList, 2)
	assert.Nil(t, err)
}

func TestTrade__UnmarshalJSON_Retruns_Nil_for_EmptyString(t *testing.T) {
	trade := kraken.Trade{}
	assert.Nil(t, trade.UnmarshalJSON(([]byte(`""`))))
	assert.Nil(t, trade.UnmarshalJSON(([]byte("null"))))
}

func TestTrade__UnmarshalJSON_Retruns_Error_For_BadData(t *testing.T) {
	trade := kraken.Trade{}
	err := trade.UnmarshalJSON(([]byte("abc : 123")))
	assert.Error(t, err)
}

func TestTrade__UnmarshalJSON_Map_Data_Correctly_For_GoodData(t *testing.T) {
	goodData := `[
		"53960.00000",
		"0.00136938",
		1726341504.0383348,
		"b",
		"l",
		"",
		89154536
	]`

	expected, _ := decimal.NewFromString("53960.00000")

	trade := kraken.Trade{}
	err := trade.UnmarshalJSON(([]byte(goodData)))

	assert.NotEmpty(t, trade)
	assert.Equal(t, expected, trade.Price())
	assert.Nil(t, err)
}
