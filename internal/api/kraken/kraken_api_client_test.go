package kraken_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sergioazevedo/argentum_go/internal/api/kraken"
	"github.com/sergioazevedo/argentum_go/internal/lib/servertest"
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
	assert.Equal(t, tradeList[0].Date, time.Unix(1726341504, 383348).UTC())
	assert.Equal(t, tradeList[1].Date, time.Unix(1726341510, 5618985).UTC())
	assert.Nil(t, err)
}
