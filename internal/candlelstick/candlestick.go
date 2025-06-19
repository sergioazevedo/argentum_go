package candlestick

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/sergioazevedo/argentum_go/internal/trade"
)

type Candlestick struct {
	Open   decimal.Decimal
	Close  decimal.Decimal
	High   decimal.Decimal
	Low    decimal.Decimal
	Volume decimal.Decimal
	Date   time.Time
}

func FromTrades(trades []trade.Trade, interval string) []Candlestick {
	maxInterval, err := time.ParseDuration(interval)
	if err != nil {
		return nil
	}
	currentDate := trades[0].Date

	candle := Candlestick{
		Date:   currentDate,
		Open:   trades[0].Price,
		High:   trades[0].Price,
		Low:    trades[0].Price,
		Volume: trades[0].Volume,
	}
	result := []Candlestick{}

	for i, v := range trades {
		tradeInteval := v.Date.Sub(currentDate)
		// check if the current trade is in candle the interval
		if tradeInteval < maxInterval {
			candle.Volume = candle.Volume.Add(v.Volume)
			// check high price
			if v.Price.GreaterThan(candle.High) {
				candle.High = v.Price
			}
			// check low price
			if v.Price.LessThan(candle.Low) {
				candle.Low = v.Price
			}
		} else {
			// the trade is out of candle interval
			// fetch previous trade price as candle closing price
			candle.Close = trades[i-1].Price
			// adds candle to the result list
			result = append(result, candle)
			//initialize the new base Date and a new candle
			currentDate = v.Date
			candle = Candlestick{
				Date: currentDate,
				Open: v.Price,
				High: v.Price,
				Low:  v.Price,
			}
		}
	}
	// fetches last trade price as candle closing price
	candle.Close = trades[len(trades)-1].Price

	return append(result, candle)
}
