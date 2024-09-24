package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/sergioazevedo/argentum_go/internal/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMustCreateATradeWithNoError(t *testing.T) {
	now := time.Now()
	trade, err := models.NewTrade(now, decimal.New(10, 0), decimal.New(10, 22))
	assert.Nil(t, err)
	assert.Equal(t, trade.Quantity(), decimal.New(10, 0))
	assert.True(t, trade.Date().Equal(now))
	assert.True(t, trade.Price().Equal(decimal.New(10, 22)))
}

func TestTradeVolume(t *testing.T) {
	trade, _ := models.NewTrade(time.Now(), decimal.New(10, 0), decimal.New(20, 0))
	assert.Equal(t, trade.Volume(), decimal.NewFromInt32(200))
}

func TestTradeCantHaveNegativeQuantity(t *testing.T) {
	trade, err := models.NewTrade(time.Now(), decimal.New(-10, 0), decimal.New(-10, 0))

	assert.Empty(t, trade)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "quantitiy must be > 0")
}

func TestTradeCantHaveNegativePrices(t *testing.T) {
	trade, err := models.NewTrade(time.Now(), decimal.New(10, 0), decimal.New(-10, 0))

	assert.Empty(t, trade)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "price must be >= 0.0")
}

func TestTradeMustHaveADate(t *testing.T) {
	trade, err := models.NewTrade(time.Time{}, decimal.New(10, 0), decimal.New(10, 22))

	assert.Empty(t, trade)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "date is missing")
}

func TestTradesFromEmptyJSON(t *testing.T) {
	jsonString := "[]"
	var jsonData models.TradesJSONData
	_ = json.Unmarshal([]byte(jsonString), &jsonData)

	tradeList := models.TradesFromJSON(jsonData)
	assert.Empty(t, tradeList)
}

func TestTradesFromJSONWithNoData(t *testing.T) {
	jsonString := "[{}]"
	var jsonData models.TradesJSONData
	_ = json.Unmarshal([]byte(jsonString), &jsonData)

	tradeList := models.TradesFromJSON(jsonData)
	assert.Empty(t, tradeList)
}

func TestTradesFromBinanceJSONData(t *testing.T) {
	jsonString := `[
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
    },
    {
        "id": 136839515,
        "price": "53878.47000000",
        "qty": "0.00619000",
        "quoteQty": "333.50772930",
        "time": 1726338644307,
        "isBuyerMaker": true,
        "isBestMatch": true
    },
    {
        "id": 136839516,
        "price": "53878.42000000",
        "qty": "0.00620000",
        "quoteQty": "334.04620400",
        "time": 1726338646307,
        "isBuyerMaker": true,
        "isBestMatch": true
    },
    {
        "id": 136839517,
        "price": "53878.52000000",
        "qty": "0.00620000",
        "quoteQty": "334.04682400",
        "time": 1726338647307,
        "isBuyerMaker": true,
        "isBestMatch": true
    },
    {
        "id": 136839518,
        "price": "53896.62000000",
        "qty": "0.00619000",
        "quoteQty": "333.62007780",
        "time": 1726338678138,
        "isBuyerMaker": false,
        "isBestMatch": true
    }
	]`

	var jsonData models.TradesJSONData
	_ = json.Unmarshal([]byte(jsonString), &jsonData)

	tradeList := models.TradesFromJSON(jsonData)
	assert.Len(t, tradeList, 6)
}
