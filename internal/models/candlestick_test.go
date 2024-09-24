package models_test

import (
	"encoding/json"
	"testing"

	"github.com/sergioazevedo/argentum_go/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCandlesticsFrom(t *testing.T) {
	trades := tradeList()

	candles := models.CadlesticksFrom(trades, "10s")
	assert.Len(t, candles, 3)
}

func tradeList() []models.Trade {
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

	return models.TradesFromJSON(jsonData)
}
