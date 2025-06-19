package candlestick_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"

	candlestick "github.com/sergioazevedo/argentum_go/internal/candlelstick"
	trade "github.com/sergioazevedo/argentum_go/internal/trade"
)

func TestCandlesticsFrom(t *testing.T) {
	trades := tradeList(10, "10m", 2, time.Time{})

	candles := candlestick.FromTrades(trades, "24h")
	assert.Len(t, candles, 5)
}

func tradeList(total int, interval string, perDay int, startDate time.Time) []trade.Trade {
	var currentDate time.Time
	dateInterval, _ := time.ParseDuration(interval)
	if startDate.IsZero() {
		currentDate = time.Now().Truncate(time.Hour)
	} else {
		currentDate = startDate
	}

	result := make([]trade.Trade, 0, total)
	totalPerDay := 0
	for i := 0; i < total; i++ {
		qty := decimal.NewFromFloat(randomFloat64(0.3, 20.0))
		price := decimal.NewFromFloat(randomFloat64(10.45, 98.87))
		volume := price.Mul(qty)

		result = append(result, trade.Trade{
			Date:     currentDate,
			Quantity: qty,
			Price:    price,
			Volume:   volume,
		})

		totalPerDay++
		currentDate = currentDate.Add(dateInterval)

		if totalPerDay == perDay {
			currentDate = currentDate.
				AddDate(0, 0, 1).
				Truncate(time.Hour)

			totalPerDay = 0
		}
	}
	return result
}

// generate random float in range of min and max inclusive
func randomFloat64(min, max float64) float64 {
	rand.Seed(102478)
	return (min + rand.Float64()*(max-min))
}
