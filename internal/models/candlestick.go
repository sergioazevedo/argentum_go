package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Candlestick struct {
	open   decimal.Decimal
	close  decimal.Decimal
	high   decimal.Decimal
	low    decimal.Decimal
	volume decimal.Decimal
	date   time.Time
}

func (c Candlestick) Open() decimal.Decimal {
	return c.open
}

func (c Candlestick) Close() decimal.Decimal {
	return c.close
}

func (c Candlestick) High() decimal.Decimal {
	return c.high
}

func (c Candlestick) Low() decimal.Decimal {
	return c.low
}

func (c Candlestick) Volume() decimal.Decimal {
	return c.volume
}

func (c Candlestick) Date() time.Time {
	return c.date
}

func CadlesticksFrom(trades []Trade, interval string) []Candlestick {
	maxInterval, _ := time.ParseDuration(interval)
	currentDate := trades[0].Date()

	candle := Candlestick{
		date:   currentDate,
		open:   trades[0].Price(),
		high:   trades[0].Price(),
		low:    trades[0].Price(),
		volume: trades[0].Volume(),
	}
	result := make([]Candlestick, 0, 50)

	for i, v := range trades {
		tradeInteval := v.Date().Sub(currentDate)
		// check if the current trade is in candle the interval
		if tradeInteval <= maxInterval {
			candle.volume = candle.volume.Add(v.Volume())
			// check high price
			if v.Price().GreaterThan(candle.high) {
				candle.high = v.Price()
			}
			// check low price
			if v.Price().LessThan(candle.low) {
				candle.low = v.Price()
			}
		} else {
			// the trade is out of candle interval
			// fetch previous trade price as candle closing price
			candle.close = trades[i-1].Price()
			// adds candle to the result list
			result = append(result, candle)
			//initialize the new base Date and a new candle
			currentDate = v.Date()
			candle = Candlestick{
				date: currentDate,
				open: v.Price(),
				high: v.Price(),
				low:  v.Price(),
			}
		}
	}
	// fetches last trade price as candle closing price
	candle.close = trades[len(trades)-1].Price()

	return append(result, candle)
}
